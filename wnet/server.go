package wnet

import (
    "fmt"
    "net"
    "tcpw/utils"
    "tcpw/wiface"
)

type Server struct {
    Name      string
    IPVersion string
    IP        string
    Port      int
    Router    wiface.IRouter
}

func (s *Server) Start() {
    fmt.Printf("[TCPw] Server name: %s,listenner at IP: %s, Port %d is starting\n",
        utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TCPPort)
    fmt.Printf("[TCPw] Version: %s,MaxConn: %d, MaxPackageSize %d\n",
        utils.GlobalObject.Version,
        utils.GlobalObject.MaxConn,
        utils.GlobalObject.MaxPackageSize)

    go func() {
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

            dealConn := NewConnection(conn, cid, s.Router)
            cid++

            go dealConn.Start()
        }
    }()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
    s.Start()

    select {}
}

func (s *Server) AddRouter(router wiface.IRouter) {
    s.Router = router
    fmt.Println("Add Router Succ!!")
}

func NewServer(name string) wiface.IServer {
    s := &Server{
        Name:      utils.GlobalObject.Name,
        IPVersion: "tcp4",
        IP:        utils.GlobalObject.Host,
        Port:      utils.GlobalObject.TCPPort,
        Router:    nil,
    }

    return s
}
