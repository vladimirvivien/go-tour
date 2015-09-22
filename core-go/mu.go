package main

import "fmt"
import "strings"

// Muting Multi-Value Returns
// In current version of Go (1.2), the multi-value return can only
// be accessed by doing a multi-variable assignment:
//
// b, c := a()
//
// That works great.  However, sometimes, instead of getting to the interested value using
// a two-step approach (assignment, then access), I want to skip the assigment
// step and do an inline access of the value.  That would allow me to chain my expressions.
// However, when your returns are multi-valued, you cannot do that in Go.  For instance
// the following will blow up:
//
// Path{FileName:path.Split("/a/b/c.d")}
//
// The Go compiler will complain (paraphrasing here) about multi-value in a single-value context.
//
// So I started to wonder, well, how is the multi-value return represented inside Go?  Is array,
// a slice, container, what ?  And why fmt.Println(path.Split(...)) works ?
// Looking for the answer to the ladder question in Go's source revealed some interesting facts
// about the multi-value function returns:
//
// 1) Multi-value returns will be cast to []interface{} when passed as function params.
//    Meaning, fmt.Printf(a...interface{}) can receive func Split()(string,string).
// 2) The compiler only does that trick (as far as I know) for function params.  Any other
//    context yield to big fat failures:
//    - a [] interface{} := path.Split(...) // illegal, fails compilation
//    - var b []interface{} = path.Split(...) // fails with multi-value in single value context
//    - var c []interface{}; c = path.Split(...) // stil fails (as above)
//
// So, that led me to create mute function (represented by greek mu, single character, simple).
//
// func µ(a ...interface{}) []interface{} {
// 	return a
// }
//
// This function does a conversion of variable params "x...interface{}"" to a standard []interface{}.
// The implication here is that you can treat the multi-value returns as a regular array and
// target any of its index as follows:
//
// Path{FileName:µ(path.Split("/a/b/c.d"))[0].(string)}
//
// NOTE: since µ() returns []interface{}, you will need Type Assertion to get the actual value.
//
// Execute the code below on Go Playground - http://play.golang.org/p/IwqmoKwVm-

func µ(a ...interface{}) []interface{} {
	return a
}

type A struct {
	B string
	C func() string
}

func main() {
	a := A{
		B: strings.TrimSpace(µ(E())[1].(string)),
		C: µ(G())[0].(func() string),
	}

	fmt.Printf("%s says %s\n", a.B, a.C())
}

func E() (bool, string) {
	return false, "F"
}

func G() (func() string, bool) {
	return func() string { return "Hello" }, true
}
