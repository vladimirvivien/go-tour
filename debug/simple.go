package main

import "fmt"
import "math/big"

// recursive factorial with big package
func f(n int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	next := big.NewInt(int64(n))
	var result big.Int
	return result.Mul(f(n-1), next)
}

func main() {
	for i := 0; i < 100; i++ {
		go fmt.Println(f(i))
	}
}
