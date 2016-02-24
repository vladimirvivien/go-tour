package main

import (
	"fmt"
	"sync"
)

// run with -race flag to see race conditions
func main() {
	var m sync.Mutex
	done := false
	total := 0

	go func() {
		m.Lock()
		done = true
		m.Unlock()
		total += 2
		total += 4
	}()

	m.Lock()
	done = false
	m.Unlock()

	m.Lock()
	fmt.Println(done)
	m.Unlock()
	fmt.Println(total)
}
