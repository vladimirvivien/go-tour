package main

import (
	"fmt"
)

var (
	N       = 100
	WORKERS = 4
)

func producer(ints ...int) <-chan int {
	fmt.Println("Producer got ", ints)
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
	// use WORKERS channels to fanout processing of consumers
	fanOutChs := make([]<-chan int, WORKERS)

	// produce N numbers
	p1 := producer(
		func() []int {
			result := make([]int, N)
			for i := 0; i < N; i++ {
				result[i] = i
			}
			return result
		}()...,
	)

	// Use workers to consume & return result
	for i := 0; i < WORKERS; i++ {
		fanOutChs[i] = consumer(p1)
	}

	for _, ch := range fanOutChs {
		for val := range ch {
			fmt.Print(val)
		}
	}
}
