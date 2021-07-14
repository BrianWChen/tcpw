package main

import (
    "fmt"
    "tcpw/wiface"
    "tcpw/wnet"
)

type PingRouter struct {
    wnet.BaseRouter
}

func (this *PingRouter) Handle(request wiface.IRequest) {
    fmt.Println("Call Router Handle...")
    fmt.Println("recv from client: msgID = ", request.GetMsgID(),
        ", data = ", string(request.GetData()))

    err := request.GetConnection().SendMsg(1, []byte("ping... ping... ping"))
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    s := wnet.NewServer("[TCPW V0.5]")

    s.AddRouter(&PingRouter{})

    s.Serve()
}
