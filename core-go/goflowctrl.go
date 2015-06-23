package main

import (
	"fmt"
)

func main() {

	// regular loop consruct
	for i := 0 ; i < 10; i++ {
		fmt.Println(i)
	}

	// equivalent to do-while
	sum := 10;
	for sum >= 0 {
		fmt.Println (sum)
		sum--
	}
	// loops forever
	counter := 0;
	for {
		if counter = counter + 20; counter < 200 { // NOTE: dec/assigment in IF-statement.
			fmt.Printf ("I am at %v\n", counter)			
		}else{
			break
		}
	}

}