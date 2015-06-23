package main

import (
	"fmt"; 
	"math"
)

// Go is not object-based.
// However, structural types such as struct (or any complex types, user-defined or otherwise)
// can be the recipient of method signatures.
type Vertex struct {
	X, Y float64
}

// type Vertex receives method Scale.
// Note the receiver is specified as a pointer to the type
// Any update to the receiver will be made to the directly (via pointer.value)
// If the receiver is specified as value (v Vertex), 
// then the function will receive a copy value of the receiver.
// To see this, change the receiver's signature from v *Vertex to v Vertex and see.
func (v *Vertex) Scale (factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

// A custom float type
type MyFloat float64

// a receiving method for the MyFloat type.
// Notice Abs receives a value not a pointer.
// It's assume that each value of MyFloat will be different copy of float64.
func (f MyFloat) Abs () float64 {
	if f < 0 {
		return float64(-1 * f) 
	}else{
		return float64(f)
	}
}

func main() {
	v := Vertex{10, 20}
	fmt.Println ("Vertex unscalled", v)
	v.Scale(5)
	fmt.Println ("Vertex.Scale(5) = ", v)
	f := MyFloat(-math.Sqrt2)
	fmt.Println ("Myfloat.Abs() ", f.Abs())
}