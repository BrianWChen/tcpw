package wiface

type IRequest interface {
    GetConnection() IConnection
    GetData() []byte
    GetMsgID() uint32
}
