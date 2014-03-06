package main

import (
	"fmt"
)

/*
 * Implement Sqrt using Newton's method
 * z = z - ((z^2 - x) / 2z)
 * Where z is an increasing value.
 */
func Sqrt (x float64) float64 {
	z := float64(1)
	//prev := float64(0)
	
	for {
		fmt.Printf ("z = %v\n", z)
		z = z - ((z*z - x) / 2*z)
	}

	return z
}


func main() {
	fmt.Println (Sqrt(4))
}