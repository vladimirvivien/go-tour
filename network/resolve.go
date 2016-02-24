package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing address to resolve")
		return
	}

	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Unable to resove:", err)
		return
	}
	fmt.Println("Resoved to :", addr)
}
