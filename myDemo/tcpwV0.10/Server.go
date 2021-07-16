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

    err := request.GetConnection().SendMsg(201, []byte("Hello Zinx Router V0.9"))
    if err != nil {
        fmt.Println(err)
    }
}

//创建连接的时候执行
func DoConnectionBegin(conn wiface.IConnection) {
    fmt.Println("DoConnecionBegin is Called ... ")
    err := conn.SendMsg(202, []byte("DoConnection BEGIN..."))
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Set conn Name, Hoe...")
    conn.SetProperty("Name", "BrianWChen")
    conn.SetProperty("Github", "https://github.com/BrianWChen")
}

//连接断开的时候执行
func DoConnectionLost(conn wiface.IConnection) {
    fmt.Println("DoConneciotnLost is Called ... ")
    fmt.Println("conn Id = ", conn.GetConnID(), " is Lost..")

    if name, err := conn.GetProperty("Name"); err == nil {
        fmt.Println("Name = ", name)
    }
}

func main() {
    s := wnet.NewServer("[TCPW V0.10]")

    s.SetOnConnStart(DoConnectionBegin)
    s.SetOnConnStop(DoConnectionLost)

    s.AddRouter(0, &PingRouter{})
    s.AddRouter(1, &HelloZinxRouter{})

    s.Serve()
}
