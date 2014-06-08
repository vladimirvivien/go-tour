package main

import (
    "net"
    "fmt"
    "strings"
)

func main() {
    fmt.Println ("Next available port:", nextAvailPort())
    fmt.Println ("Another available port:", nextAvailPort())
}

func nextAvailPort() string {
    l, _ := net.Listen("tcp", ":0")
    defer l.Close()
    addr := l.Addr().String() 
    return addr[strings.LastIndex(addr, ":")+1:]
}
