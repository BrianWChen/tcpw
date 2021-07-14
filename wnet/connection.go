package wnet

import (
    "errors"
    "fmt"
    "io"
    "net"
    "tcpw/wiface"
)

type Connection struct {
    Conn      *net.TCPConn
    ConnID    uint32
    isClosed  bool
    handleAPI wiface.HandleFunc
    ExitChan  chan bool
    Router    wiface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router wiface.IRouter) *Connection {
    c := &Connection{
        Conn:     conn,
        ConnID:   connID,
        Router:   router,
        isClosed: false,
        ExitChan: make(chan bool, 1),
    }

    return c
}

func (c *Connection) StartReader() {
    fmt.Println("Reader Goroutine is running...")
    defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
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

        go func(request wiface.IRequest) {
           c.Router.PreHandle(request)
           c.Router.Handle(request)
           c.Router.PostHandle(request)
        }(&req)

    }
}

func (c *Connection) Start() {
    fmt.Println("Conn Start()... ConnId = ", c.ConnID)
    go c.StartReader()
}

func (c *Connection) Stop() {
    fmt.Println("Conn Stop()... ConnId = ", c.ConnID)

    if c.isClosed == true {
        return
    }

    c.isClosed = true

    c.Conn.Close()

    close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
    return c.Conn
}

func (c *Connection) GETConnID() uint32 {
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
    if _, err := c.Conn.Write(msg); err != nil {
        fmt.Println("Write msg id ", msgID, "error :", err)
        return errors.New("conn Write error")
    }

    return nil
}
