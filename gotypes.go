package main

import (
	"fmt" /*"math/cmplx"*/)

const FMT = "%T Value = %v\n"

var (
	strVal  string = "String val"
	big64   uint64 = 1<<64 - 1
	intVal  int    = 1<<32 - 1
	runeVal rune   = 'C'
)

func main() {
	fmt.Printf(FMT, big64, big64)
	fmt.Printf(FMT, intVal, intVal)
	fmt.Printf(FMT, strVal, strVal)
	fmt.Printf(FMT, runeVal, runeVal)
}
