package main

import "fmt"

func main() {
	// a slice is a *(pointer)Type to an array
	// literal form of a slice of type []T
	p := []int{12, 44, 5, 66, 100, 524, 52}

	fmt.Printf("slice p = %v (has type %T) \n", p, p)

	// lenght function takes a slice param
	fmt.Printf("len(p) = %d, cap(p)=%d\n", len(p), cap(p))

	// re-slicing slice p returns new slice
	p1 := p[1:5]
	fmt.Printf("re-slicing slice p[1:5] = %v\n", p1)

	// slice s[:N] returns new slice starting from index 0
	p2 := p[:4]
	fmt.Printf("re-slicing slice p[0:4] = %v\n", p2)

	// slice s[i:] returns new slice starting from i to len(s)
	fmt.Printf("re-slicing slice p[3:] = %v\n", p[3:])

	// non literal creation uses function 'make()'
	m := make([]int, 5, 10) // creates a new slice []T
	fmt.Printf("slice m = %v (has type %T) \n", m, m)
	fmt.Printf("len(m) = %d, cap(m)=%d\n", len(m), cap(m))

	// ************ Range of Slices ***************//
	// range keyword reurns and index i, and value v from a slice
	for i, v := range p2 {
		fmt.Printf("i=%d \t v=%d\n", i, v)
	}

	// if you want just the index, omit the value
	for i := range p1 {
		fmt.Printf("i=%d\n", i)
	}
	// conversely, if you want value only, mute the index as so
	for _, v := range p2 {
		fmt.Printf("v=%d\n", v)
	}

}
