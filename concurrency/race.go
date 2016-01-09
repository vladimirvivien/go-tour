package main

import "fmt"

func main() {
	state := 0
	wait := make(chan struct{})

	go func() {
		total := 0
		for i := 0; i < 12; i++ {
			total = 2*i + total
		}
		state = 1
		close(wait)
	}()

	if state == 1 {
		fmt.Println(state)
	}
	<-wait
}
