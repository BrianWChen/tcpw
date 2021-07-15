package wnet

import (
    "fmt"
    "net"
    "tcpw/utils"
    "tcpw/wiface"
)

type Server struct {
    Name       string
    IPVersion  string
    IP         string
    Port       int
    MsgHandler wiface.IMsgHandler
    ConnMgr    wiface.IConnManager
    OnConnStart func(conn wiface.IConnection)
    OnConnStop func(conn wiface.IConnection)
}

func (s *Server) Start() {
    fmt.Printf("[TCPw] Server name: %s,listenner at IP: %s, Port %d is starting\n",
        utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TCPPort)
    fmt.Printf("[TCPw] Version: %s,MaxConn: %d, MaxPackageSize %d\n",
        utils.GlobalObject.Version,
        utils.GlobalObject.MaxConn,
        utils.GlobalObject.MaxPackageSize)

    go func() {
        //0 启动worker工作池机制
        s.MsgHandler.StartWorkerPool()

        addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
        if err != nil {
            fmt.Println("resolve tcp addr error: ", err)
            return
        }

        listener, err := net.ListenTCP(s.IPVersion, addr)
        if err != nil {
            fmt.Println("listen ", s.IPVersion, " err ", err)
            return
        }

        fmt.Println("start TCPW server succ ", s.Name, " succ, Listening... ")
        var cid uint32
        cid = 0

        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Println("Accept err ", err)
                continue
            }

            if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
                fmt.Println("Too Many Connections MaxConn = ", utils.GlobalObject.MaxConn)
                conn.Close()
                continue
            }

            dealConn := NewConnection(s, conn, cid, s.MsgHandler)
            cid++

            go dealConn.Start()
        }
    }()
}

func (s *Server) Stop() {
    fmt.Println("[STOP] TCPw server name ", s.Name)
    s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
    s.Start()

    select {}
}

func (s *Server) AddRouter(msgID uint32, router wiface.IRouter) {
    s.MsgHandler.AddRouter(msgID, router)
    fmt.Println("Add Router Succ!!")
}

func (s *Server) GetConnMgr() wiface.IConnManager {
    return s.ConnMgr
}

//SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(wiface.IConnection)) {
    s.OnConnStart = hookFunc
}

//SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(wiface.IConnection)) {
    s.OnConnStop = hookFunc
}

//CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn wiface.IConnection) {
    if s.OnConnStart != nil {
        fmt.Println("---> CallOnConnStart....")
        s.OnConnStart(conn)
    }
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn wiface.IConnection) {
    if s.OnConnStop != nil {
        fmt.Println("---> CallOnConnStop....")
        s.OnConnStop(conn)
    }
}

func NewServer(name string) wiface.IServer {
    s := &Server{
        Name:       utils.GlobalObject.Name,
        IPVersion:  "tcp4",
        IP:         utils.GlobalObject.Host,
        Port:       utils.GlobalObject.TCPPort,
        MsgHandler: NewMsgHandle(),
        ConnMgr:    NewConnManager(),
    }

    return s
}
