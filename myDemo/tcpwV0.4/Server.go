package main

import (
    "fmt"
    "tcpw/wiface"
    "tcpw/wnet"
)

type PingRouter struct {
    wnet.BaseRouter
}

func (this *PingRouter) PreHandle(request wiface.IRequest) {
    fmt.Println("Call Router PreHandle...")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
    if err != nil {
        fmt.Println("call back before ping error")
    }
}

func (this *PingRouter) Handle(request wiface.IRequest) {
    fmt.Println("Call Router Handle...")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping... ping...\n"))
    if err != nil {
        fmt.Println("call back ping... ping... ping... error")
    }
}

func (this *PingRouter) PostHandle(request wiface.IRequest) {
    fmt.Println("Call Router PostHandle...")
    _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
    if err != nil {
        fmt.Println("call back after ping error")
    }
}

func main() {
    s := wnet.NewServer("[TCPW V0.4]")

    s.AddRouter(&PingRouter{})

    s.Serve()
}
