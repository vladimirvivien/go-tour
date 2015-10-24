package main

import "fmt"
import "time"

var evCh chan *Act
var blockCh chan struct{}

type Act struct {
	name  string
	param interface{}
}

func main() {
	evCh = make(chan *Act, 1024)
	go evLoop()
	blockCh = make(chan struct{})
	fmt.Println("Starting...")
	go Action("W")

	time.Sleep(time.Millisecond * 1000)
	close(blockCh)
	time.Sleep(time.Millisecond * 1000)

}

func evLoop() {
	fmt.Println("Starting ev Loop")
	for {
		select {
		case a := <-evCh:
			invoke(a)
		}
	}
}

func Action(s string) {
	ch := make(chan struct{})
	evCh <- &Act{name: s, param: ch}
	<-ch
	fmt.Println("Finished Action ...")
}

func invoke(a *Act) {
	ch := a.param.(chan struct{})
	time.Sleep(time.Millisecond * 200)
	<-blockCh
	close(ch)
}
