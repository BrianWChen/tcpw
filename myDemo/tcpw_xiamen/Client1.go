package main

import (
    "fmt"
    "net"
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
        //_, err0 := conn.Write([]byte("CTRL:10;666666#Z#"))
        //if err0 != nil {
        //    fmt.Println("writ conn err ", err)
        //    return
        //}
        _, err1 := conn.Write([]byte("ST:10;Product:Product;Hardware:V2.0;Software:V2.0"))
        if err1 != nil {
            fmt.Println("writ conn err ", err)
            return
        }

        _, err2 := conn.Write([]byte("ST:19;Product:Product;Hardware:V2.0;Software:V2.0"))
        if err2 != nil {
            fmt.Println("writ conn err ", err)
            return
        }

        time.Sleep(30 * time.Millisecond)
    }
}
