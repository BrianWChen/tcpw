package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    fmt.Println("client start...")

    time.Sleep(1 * time.Second)

    conn, err := net.Dial("tcp", "127.0.0.1:8999")
    if err != nil {
        fmt.Println("client start err, exit!")
        return
    }

    for {
        _, err := conn.Write([]byte("Hello TCPw V0.3"))
        if err != nil {
            fmt.Println("writ conn err ", err)
            return
        }

        buf := make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read buf err ", err)
            return
        }

        fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

        time.Sleep(1 * time.Second)
    }
}
