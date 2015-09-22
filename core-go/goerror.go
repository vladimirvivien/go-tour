package main

import (
	"fmt"
)

// Go includes the built-in error interface defined as
/*
    type error interface {
		Error () string
    }

    So, any value that  satisfies this interface can be used wherever errors are used.
*/
//

type GreetingError struct {
	Who string
}

func (ge *GreetingError) Error() string {
	return fmt.Sprintf("Tried to say hello to %v, but cant.", ge.Who)
}

func SayHello(who string) (ok bool, err *GreetingError) {
	switch {
	case who == "World":
		ok = true
	case who == "world":
		ok = true
	default:
		ok = false
		err = &GreetingError{who}
	}
	return ok, err
}

func main() {
	if ok, err := SayHello("Moon"); ok == true {
		fmt.Println("Said hello, error ", err)
	} else {
		fmt.Println(err.Error())
	}
}
