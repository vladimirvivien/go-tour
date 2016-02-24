package main

import "fmt"

var slice []byte

func main() {
	slice = make([]byte, 12)
	go func() {
		slice[0] = byte('H')
	}()
	slice[3] = byte('e')
	fmt.Println(slice)
}
