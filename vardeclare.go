package main

import "fmt"

// var dec outside of func body
// must have dec keyword var
// no shortcut := assigment operator allowed
var a, b, c = 1, "me", false
var d, e, f int

// var block
var (
	k rune = 1234
	l int  = 122
	m      = false
)

// constants
const str = "This will never change"

func main() {
	fmt.Println(a, b, c)
	fmt.Println("Other values ")

	d, e, f = 10, 12, 44
	fmt.Println(d, e, f)

	// dec shortcut allowed here
	g, h := "Mars", "Rover"
	fmt.Println(g, h)

	// print dec block
	fmt.Println(k, l, m)

	fmt.Println(str)
}
