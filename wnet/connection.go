package wnet

import (
    "errors"
    "fmt"
    "io"
    "net"
    "sync"
    "tcpw/utils"
    "tcpw/wiface"
)

type Connection struct {
    TcpServer  wiface.IServer
    Conn       *net.TCPConn
    ConnID     uint32
    isClosed   bool
    handleAPI  wiface.HandleFunc
    ExitChan   chan bool
    msgChan    chan []byte
    MsgHandler wiface.IMsgHandler
    //链接属性
    property map[string]interface{}
    ////保护当前property的锁
    propertyLock sync.Mutex
}

func NewConnection(server wiface.IServer, conn *net.TCPConn, connID uint32, msgHandler wiface.IMsgHandler) *Connection {
    c := &Connection{
        TcpServer:  server,
        Conn:       conn,
        ConnID:     connID,
        MsgHandler: msgHandler,
        isClosed:   false,
        msgChan:    make(chan []byte),
        ExitChan:   make(chan bool, 1),
        property:   make(map[string]interface{}),
    }

    c.TcpServer.GetConnMgr().Add(c)

    return c
}

func (c *Connection) StartReader() {
    fmt.Println("Reader Goroutine is running...")
    defer fmt.Println("[Reader is exit!] connID = ", c.ConnID, " remote addr is ", c.RemoteAddr().String())
    defer c.Stop()

    for {
        //buf := make([]byte, utils.GlobalObject.MaxPackageSize)
        //cnt, err := c.Conn.Read(buf)
        //if err != nil {
        //    fmt.Println("recv buf err ", err)
        //    continue
        //}

        //fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

        dp := NewDataPack()

        headData := make([]byte, dp.GetHeadLen())
        if _, err := io.ReadFull(c.Conn, headData); err != nil {
            fmt.Println("read msg head error ", err)
            return
        }
        //fmt.Printf("read headData %+v\n", headData)

        //拆包，得到msgID 和 datalen 放在msg中
        msg, err := dp.Unpack(headData)
        if err != nil {
            fmt.Println("unpack error ", err)
            return
        }

        //根据 dataLen 读取 data，放在msg.Data中
        var data []byte
        if msg.GetDataLen() > 0 {
            data = make([]byte, msg.GetDataLen())
            if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
                fmt.Println("read msg data error ", err)
                return
            }
        }
        msg.SetData(data)

        //得到当前客户端请求的Request数据
        req := Request{
            conn: c,
            msg:  msg,
        }

        if utils.GlobalObject.WorkerPoolSize > 0 {
            //已经启动工作池机制，将消息交给Worker处理
            c.MsgHandler.SendMsgToTaskQueue(&req)
        } else {
            //从绑定好的消息和对应的处理方法中执行对应的Handle方法
            go c.MsgHandler.DoMsgHandler(&req)
        }
    }
}

func (c *Connection) StartWriter() {
    fmt.Println("[Writer Goroutine is running]")
    defer fmt.Println("[conn Writer exit!]", c.RemoteAddr().String())

    for {
        select {
        case data := <-c.msgChan:
            //有数据要写给客户端
            if _, err := c.Conn.Write(data); err != nil {
                fmt.Println("Send Data error:, ", err, " Conn Writer exit")
                return
            }
            //fmt.Printf("Send data succ! data = %+v\n", data)
        case <-c.ExitChan:
            return
        }
    }
}

func (c *Connection) Start() {
    fmt.Println("Conn Start()... ConnId = ", c.ConnID)
    go c.StartReader()
    go c.StartWriter()

    c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
    fmt.Println("Conn Stop()... ConnId = ", c.ConnID)

    if c.isClosed == true {
        return
    }

    c.isClosed = true

    c.TcpServer.CallOnConnStop(c)

    c.Conn.Close()

    c.ExitChan <- true

    c.TcpServer.GetConnMgr().Remove(c)

    close(c.ExitChan)
    close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
    return c.Conn
}

func (c *Connection) GetConnID() uint32 {
    return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
    return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
    if c.isClosed == true {
        return errors.New("connection closed when send msg")
    }

    //将data封包，并且发送
    dp := NewDataPack()
    msg, err := dp.Pack(NewMsgPackage(msgID, data))
    if err != nil {
        fmt.Println("Pack error msg ID = ", msgID)
        return errors.New("Pack error msg ")
    }

    //写回客户端
    c.msgChan <- msg

    return nil
}

//SetProperty 设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
    c.propertyLock.Lock()
    defer c.propertyLock.Unlock()
    if c.property == nil {
        c.property = make(map[string]interface{})
    }

    c.property[key] = value
}

//GetProperty 获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
    c.propertyLock.Lock()
    defer c.propertyLock.Unlock()

    if value, ok := c.property[key]; ok {
        return value, nil
    }

    return nil, errors.New("no property found")
}

//RemoveProperty 移除链接属性
func (c *Connection) RemoveProperty(key string) {
    c.propertyLock.Lock()
    defer c.propertyLock.Unlock()

    delete(c.property, key)
}
