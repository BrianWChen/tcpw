package wiface

import "net"

type IConnection interface {
    Start()
    Stop()
    GetTCPConnection() *net.TCPConn
    GetConnID() uint32
    RemoteAddr() net.Addr
    SendMsg(msgID uint32, data []byte) error

    SetProperty(key string, value interface{})   //设置链接属性
    GetProperty(key string) (interface{}, error) //获取链接属性
    RemoveProperty(key string)                   //移除链接属性
}

type HandleFunc func(*net.TCPConn, []byte, int) error
