package wnet

import "tcpw/wiface"

type Request struct {
    conn wiface.IConnection
    data []byte
}

func (r *Request) GetConnection() wiface.IConnection {
    return r.conn
}

func (r *Request) GetData() []byte {
    return r.data
}
