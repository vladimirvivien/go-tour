package main

import "fmt"

// NOTE: this is not a real way to do things.
// It's all illustrative.

// fibonacci is a function that returns
// a function that returns an int.

// it uses the function closer as well as the recursive implementation of fibonacci.

func fibonacci() func(i int) int {

	return func(i int) int {
		if i <= 1 {
			return i
		} else {
			// fibonacci() returns a functiona that returns the closure we need.
			return fibonacci()(i-1) + fibonacci()(i-2)
		}
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f(i))
	}
}
