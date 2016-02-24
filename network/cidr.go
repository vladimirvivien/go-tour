package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing IP address argument")
		return
	}
	cidr := os.Args[1]
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Unable to parse CIDR address", err)
		return
	}

	fmt.Println("IP address: ", ip)
	fmt.Println("Network address:", ipnet.String())

}
