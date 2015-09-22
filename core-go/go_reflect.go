package main

import (
	"fmt"
	"reflect"
)

/*
	The reflect package contains all method to introspect types in Go.
	All Types implicitly implement the empty interface type interface{}
	As such, all types (be it simple or complex) is a descendent of interface{}
*/

func main() {
	var f float64
	f = 44.5

	// reflection can occur on any value since all values is of type interface{}
	t := reflect.TypeOf(f)  // returns type information
	v := reflect.ValueOf(f) // returns concrete value info about type

	fmt.Println("Variable f of type ", t, " with value ", v.Float())

	// v.Interface() returns the concrete type as an interface{}
	y := v.Interface()
	fmt.Println("Variable y is of type ", reflect.TypeOf(y), " with value ", y)

	// Settability - the ability to update a value via reflection
	// Settability behaves the same way as does Go in other context
	// where values are passed around:
	// 1. if you reflect on a copy of a value, you will not b able to update it.
	//    The system will panic.
	// 2. To receive settability, reflection mus be done on the address of a value
	// 3. Use the idiom below to update
	//
	// NOTE: v.SetFloat(23.0) will panic

	p := reflect.ValueOf(&f)
	pp := p.Elem() // a pointer to the pointer of p (&f)
	pp.SetFloat(4.45)
	fmt.Println("Updating pointer to f of type ", pp.Kind(), " with value ", pp.Interface())
	fmt.Println("Variable f now has value ", f)

	// Complex Types (Struct, etc)
	// The reflection rules is similar for complex/custom types.
	// 1. You can navigate freely over the fields and methods
	// 2. You can get access only exported fields and methods
	// 3. You can update values / make calls to exported methods

	type S struct {
		A int
		B string
		c float64
	}

	s := S{A: 12, B: "Id", c: 3.314}
	ps := reflect.ValueOf(&s).Elem() // pointer to the pointer to s

	// walk fields/methods
	for j := 0; j < ps.NumField(); j++ {
		fi := ps.Field(j)
		if fi.CanInterface() {
			fmt.Printf("%s = %v\n", ps.Type().Field(j).Name, fi.Interface())
		} else {
			fmt.Printf("Unable to access Field %v \n", ps.Type().Field(j).Name)
		}
	}
}
