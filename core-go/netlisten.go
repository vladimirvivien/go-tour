package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", net.JoinHostPort("", ""))
	if err != nil {
		fmt.Println("Error, listening: ", err)
	}

	fmt.Println("Listening ", ln.Addr().String())
}
