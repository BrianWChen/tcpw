package main

import (
    "fmt"
    "io"
    "net"
    "tcpw/wnet"
    "time"
)

func main() {
    fmt.Println("client1 start...")

    time.Sleep(1 * time.Second)

    conn, err := net.Dial("tcp", "127.0.0.1:5000")
    if err != nil {
        fmt.Println("client start err, exit!")
        return
    }

    for {
        //发封包message消息
        dp := wnet.NewDataPack()
        msg, _ := dp.Pack(wnet.NewMsgPackage(1, []byte("TCPw V0.7 Client1 Test Message")))
        _, err := conn.Write(msg)
        if err != nil {
            fmt.Println("write error err ", err)
            return
        }

        //先读出流中的head部分
        headData := make([]byte, dp.GetHeadLen())
        _, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
        if err != nil {
            fmt.Println("read head error")
            break
        }
        //将headData字节流 拆包到msg中
        msgHead, err := dp.Unpack(headData)
        if err != nil {
            fmt.Println("server unpack err:", err)
            return
        }

        if msgHead.GetDataLen() > 0 {
            //msg 是有data数据的，需要再次读取data数据
            msg := msgHead.(*wnet.Message)
            msg.Data = make([]byte, msg.GetDataLen())

            //根据dataLen从io中读取字节流
            _, err := io.ReadFull(conn, msg.Data)
            if err != nil {
                fmt.Println("server unpack data err:", err)
                return
            }

            fmt.Println("==> Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
        }

        time.Sleep(1 * time.Second)
    }
}
