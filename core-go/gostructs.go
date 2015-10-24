package main

import (
	"fmt"
)

const FMT = "%T Value = %v\n"

// a struct is a collection of fields
// a struct is always a type.
type Astruct struct {
	X int
	Y int
}

func main() {
	// literal struct declaration with named variables
	// v receives type Astruct
	var v = Astruct{Y: 2}
	v.X = 7
	fmt.Printf(FMT, v, v)

	// literals declaration with all fields initialized
	//var b = Astruct{1,2}

	// Use the new() function to create the struct
	// x receives a pointer of type *Astruct
	x := new(Astruct) // equivalent to => 'var x *Abstract = new A(abstract)'
	fmt.Printf(FMT, x, x)
	x.Y = 12
	fmt.Println(&x)
}
