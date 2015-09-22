package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Next available port:", nextAvailPort())
	fmt.Println("Another available port:", nextAvailPort())
}

func nextAvailPort() int {
	l, _ := net.Listen("tcp4", ":0")
	defer l.Close()
	addr := l.Addr().String()
	port, _ := strconv.Atoi(addr[strings.LastIndex(addr, ":")+1:])
	return port
}
