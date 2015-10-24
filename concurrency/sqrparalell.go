package main

import "fmt"

func sqrSvc(done <-chan struct{}, vals <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case v := <-vals:
				out <- (v * v)
			case <-done:
				return
			}
		}
	}()

	return out
}

func main() {
	nums := make(chan int)

	// produce
	go func() {
		for i := 0; i < 100; i++ {
			nums <- i
		}
	}()

	// call sqr service
	done := make(chan struct{})
	wait := make(chan struct{})
	sqrs := sqrSvc(done, nums)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-sqrs, " ")
		}
		close(done)
		close(wait)
	}()
	<-wait
}
