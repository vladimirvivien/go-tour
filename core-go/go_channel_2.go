package main

import (
	"fmt"
	"time"
)

type Work struct {
	Job string
	Do  func()
}

const (
	CTRL_OK   = 0
	CTLR_QUIT = 99
)

func makeWorkQ(size int) chan Work {
	return make(chan Work, size)
}

func doWork(workQ chan Work, ctrl chan int) {
	for {
		select {
		case w := <-workQ:
			fmt.Println("Executing job ", w.Job)
			w.Do()
		case i := <-ctrl:
			if i == CTLR_QUIT {
				close(workQ)
				return
			}
		}
	}
}

func main() {
	workQ := makeWorkQ(5)
	ctrl := make(chan int)

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
		ctrl <- CTLR_QUIT
	}()

	// do some work
	go doWork(workQ, ctrl)
}
