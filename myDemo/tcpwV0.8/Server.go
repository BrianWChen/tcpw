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

    err := request.GetConnection().SendMsg(200, []byte("ping... ping... ping"))
    if err != nil {
        fmt.Println(err)
    }
}

type HelloZinxRouter struct {
    wnet.BaseRouter
}

//HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request wiface.IRequest) {
    fmt.Println("Call HelloZinxRouter Handle")
    //先读取客户端的数据，再回写ping...ping...ping
    fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

    err := request.GetConnection().SendMsg(201, []byte("Hello Zinx Router V0.6"))
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    s := wnet.NewServer("[TCPW V0.8]")

    s.AddRouter(0, &PingRouter{})
    s.AddRouter(1, &HelloZinxRouter{})

    s.Serve()
}
