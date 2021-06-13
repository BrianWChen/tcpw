package wnet

import (
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

        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Println("Accept err ", err)
                continue
            }

            go func() {
                for {
                    buf := make([]byte, 512)
                    cnt, err := conn.Read(buf)
                    if err != nil {
                        fmt.Println("recv buf err ", err)
                        continue
                    }

                    fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)

                    if _, err := conn.Write(buf[:cnt]); err != nil {
                        fmt.Println("write back buf err ", err)
                    }
                }
            }()
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
