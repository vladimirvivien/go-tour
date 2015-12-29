package main

import "fmt"

func main() {
    var ch chan int

    for i := range ch {
        fmt.Println(i)
    }
}
