package main

import (
	"fmt"
	"net"
)

func main() {
	const name = "127.0.0.1"
	addr := net.ParseIP(name)
	fmt.Printf("%T, %v", addr, addr)
}
