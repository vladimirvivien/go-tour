package main

import "fmt"
import "time"

type A struct {
	name   string
	respCh chan string
}

var evBus chan *A

func loop() {
	for {
		select {
		case a := <-evBus:
			rsp := handleEv(a)
			fmt.Println("Got Resp", rsp)
			a.respCh <- rsp
		default:
		}
	}
}

func main() {
	evBus = make(chan *A, 1)
	go loop()
	for i := 65; i <= 90; i++ {
		Action(string(i))
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Millisecond * 100)
}

func Action(a string) {
	act := &A{name: a, respCh: make(chan string)}
	evBus <- act
	fmt.Println("Waiting for response")
	rsp := <-act.respCh
	fmt.Println("Response", rsp)
}

func handleEv(a *A) string {
	fmt.Println("Got Action", a)
	return a.name + "_" + a.name
}
