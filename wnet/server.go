package wnet

import (
    "errors"
    "fmt"
    "net"
    "tcpw/wiface"
)

type Server struct {
    Name      string
    IPVersion string
    IP        string
    Port      int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
    fmt.Println("[Conn Handle] CallbackToClient...")

    if _, err := conn.Write(data[:cnt]); err != nil {
        fmt.Println("write back buf err ", err)
        return errors.New("CallBackToClient error")
    }

    return nil
}

func (s *Server) Start() {
    fmt.Printf("[Start] Server Listener at IP %s, Port %d, is starting\n", s.IP, s.Port)

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

            dealConn := NewConnection(conn, cid, CallBackToClient)
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

func NewServer(name string) wiface.IServer {
    s := &Server{
        Name:      name,
        IPVersion: "tcp4",
        IP:        "0.0.0.0",
        Port:      8999,
    }

    return s
}
