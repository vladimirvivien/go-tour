package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := "hellolife"
	epwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), 3)
	fmt.Println("Pwd", pwd)
	fmt.Println("epwd", string(epwd))

	badPwd := "badpwd"
	err := bcrypt.CompareHashAndPassword(epwd, []byte(badPwd))
	fmt.Println("pwd compared err", err)
}
