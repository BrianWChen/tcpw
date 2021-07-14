package wiface

type IDataPack interface {
    GetHeadLen() uint32
    Pack(msg IMessage) ([]byte, error)
    Unpack([]byte) (IMessage, error)
}
