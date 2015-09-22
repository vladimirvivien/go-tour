package main

import (
	"fmt"
	"net"
	_ "strings"
)

func main() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		switch addr.(type) {
		case *net.IPNet:
			ip := addr.(*net.IPNet)
			if !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				fmt.Printf("length %d\n", len(ip.IP))
				//fmt.Println("->", ip.String()[:strings.LastIndex(ip.String(),"/")])
				fmt.Println("->", ip.String())
			}
		}
	}
}
