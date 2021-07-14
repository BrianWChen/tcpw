package wnet

import "tcpw/wiface"

type Request struct {
    conn wiface.IConnection
    msg wiface.IMessage
}

func (r *Request) GetConnection() wiface.IConnection {
    return r.conn
}

func (r *Request) GetData() []byte {
    return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
    return r.msg.GetMsgID()
}
