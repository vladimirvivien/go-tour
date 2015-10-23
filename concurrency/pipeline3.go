package main

import "fmt"

var (
	N       = 100
	WORKERS = 4
)

func producer(ints ...int) <-chan int {
	fmt.Println(ints)
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, v := range ints {
			outCh <- v
		}
	}()
	return outCh
}

func consumer(done <-chan struct{}, input <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for v := range input {
			select {
			case outCh <- v * v:
			case <-done:
				return
			}
		}
	}()
	return outCh
}

func main() {
	// use WORKERS channels to fanout processing of consumers
	fanOutChs := make([]<-chan int, WORKERS)

	// create a collector channel for all fanout workers
	fanInCh := make(chan int, 2016)

	// signal
	done := make(chan struct{})
	defer close(done)

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

	// fan-in the multiple producer channels and  process content
	// loop over the producer channels for up to n WORKERS
	// Use coordination to wait for each worker to be done
	for i := 0; i < WORKERS; i++ {
		fanOutChs[i] = consumer(done, p1)
	}

	collect := func(ch <-chan int) {
		for val := range ch {
			select {
			case fanInCh <- val:
			}
		}
	}

	for _, ch := range fanOutChs {
		collect(ch)
	}

	for val := range fanInCh {
		fmt.Printf("%d ", val)
	}
	//close(fanInCh)
}
