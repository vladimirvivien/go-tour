package main

import ("fmt"; "math")

// Interface T is a type that can have a collection of function signatures;
// A function can be a receiver type T (as with any complex type);
// Any value of type that implements (receivers) any one method defined in T
// is defined as type T.

// The following interface defines type Absolutist with function Abs().
// Any type that receives function Abs() will implicitly be of type Absolutist.
// see below.
type Absolutist interface {
	Abs() float64
}

// A structure that can hold type64
type Vertex struct {
	X, Y float64
}
// The following makes Vertex an Absolutist
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// this causes MyFloat to be an Absolutist
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0{
		return float64(-f)
	}else{
		return float64(f)
	}
}

// this function takes an Abosolutist param.
// Any type that can be implied to be an Asolutist will work.
func PrintAbs(abs Absolutist) {
	fmt.Printf ("Abs from type %T = %v\n", abs, abs.Abs())
}

func main() {
	vert := &Vertex{4,12} // NOTE *Vertex is Absolutist, not Vertex.
	PrintAbs(vert)
	PrintAbs(MyFloat(-math.Sqrt2 * 100))
}