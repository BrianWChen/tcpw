package wnet

import (
    "bytes"
    "encoding/binary"
    "errors"
    "tcpw/utils"

    //"errors"

    //"tcpw/utils"
    "tcpw/wiface"
)

var defaultHeaderLen uint32 = 8

type DataPack struct{}

func NewDataPack() *DataPack {
    return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
    return defaultHeaderLen
}

func (dp *DataPack) Pack(msg wiface.IMessage) ([]byte, error) {
    dataBuff := bytes.NewBuffer([]byte{})

    if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
        return nil, err
    }

    if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
        return nil, err
    }

    if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
        return nil, err
    }

    return dataBuff.Bytes(), nil
}

func (dp *DataPack) Unpack(binaryData []byte) (wiface.IMessage, error) {
    dataBuff := bytes.NewReader(binaryData)

    msg := &Message{}

    if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
        return nil, err
    }

    if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
        return nil, err
    }

    if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
       return nil, errors.New("too large msg data received")
    }

    return msg, nil
}
