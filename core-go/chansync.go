// example of sync via channel
package main

import "fmt"
import "time"
import "sync"

var currVal int
var mu = &sync.Mutex{}

// bad with race conditions
// func main() {

// 	go func() {
// 		for {
// 			if (currVal % 10) == 0 {
// 				fmt.Println("Reached 10")
// 			}
// 			time.Sleep(time.Millisecond * 100)
// 		}
// 	}()

// 	for i := 0; i <= 100; i++ {
// 		currVal = i
// 		time.Sleep(time.Millisecond * 10)
// 	}

// }

// no race conditions since value is updated via communication
// not by memory access.
var valCh = make(chan int, 1024)

func main() {

	go func() {
		for {
			select {

			case v := <-valCh:
				if (v % 2) == 0 {
					mu.Lock()
					currVal = v
					mu.Unlock()
				}
			}
		}
	}()

	for i := 0; i <= 100; i++ {
		//currVal = i
		valCh <- i
		time.Sleep(time.Millisecond * 50)
		mu.Lock()
		fmt.Println("Value", currVal)
		mu.Unlock()
	}

}
