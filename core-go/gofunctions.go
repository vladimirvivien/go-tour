package main

import "fmt"

func main() {
	// in-line function def
	// Notice the declaration for of function type
	// func (param list) <return type of func> {}
	sqr := func(x int) int {
		return x * x
	}

	fmt.Println("Sqr of (", 2, ") = ", sqr(2))

	// function type assignment
	myadder := adderType()
	fmt.Println("5 added to 9 = ", myadder(5, 9))

	// using f as parameter
	calc(12, 9, func(a, b int) int { return a + b })
}

// a function that returns a function type
func adderType() func(a, b int) int {
	adder := func(a, b int) int {
		return a + b
	}
	return adder
}

func calc(a, b int, f func(x, y int) int) {
	fmt.Printf("Calculating %v	 for %d and %d = %d ", f, a, b, f(a, b))
}

// a function that returns a function as
