package main

import (
	"fmt"
)

func producer(ints ...int) <-chan int {
	outCh := make(chan int, 8)
	go func() {
		defer close(outCh)
		for _, v := range ints {
			outCh <- v
		}
	}()
	return outCh
}

func consumer(input <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for v := range input {
			select {
			case outCh <- v * v:
			}
		}
	}()
	return outCh
}

func main() {
	p1 := producer(1, 2, 3)
	c1 := consumer(p1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
}
