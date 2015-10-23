package main

import (
	"fmt"
	"time"
)

type Work struct {
	Job string
	Do  func()
}

func makeWorkQ(size int) chan Work {
	return make(chan Work, size)
}

func main() {
	workQ := makeWorkQ(5)

	// send work
	go func() {
		for i := 0; i < 4; i++ {
			w := Work{
				Job: "Print Greeting",
				Do: func() {
					fmt.Println("Hello, it's ", time.Now())
				},
			}

			workQ <- w
		}
		close(workQ)
	}()

	// do some work
	for w := range workQ {
		fmt.Println("Executing job ", w.Job)
		w.Do()
	}
}
