package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.JoinHostPort("0.0.0.0","0")
	fmt.Println ("About to listen on ", addr)
	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Println("Error, listening: ", err)
		return
	}

	fmt.Println("Listening ", ln.Addr().String())
}
