package main

import "tcpw/wnet"

func main() {
    s := wnet.NewServer("[TCPW V0.1]")

    s.Serve()
}
