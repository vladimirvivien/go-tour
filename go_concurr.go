package main

import ("fmt"; "time")

func sayGreetingManyTimes(s string, count int)  {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println (s)
	}
}

func main() {
	go sayGreetingManyTimes("Hello", 10)
	sayGreetingManyTimes(" World ", 1)
}