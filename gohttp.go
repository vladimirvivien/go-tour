package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Greeting struct {
	Greeting string
	Who      string
}

func (g Greeting) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	str := fmt.Sprintf("Sending warm %v to %v", g.Greeting, g.Who)
	fmt.Fprint(rsp, str)
}

func main() {
	http.Handle("/string", String("Hi. I handle path /string"))
	http.Handle("/greeting", Greeting{"Hello", "Vladimir"})

	// in-line handler registration
	http.HandleFunc("/call", func(rsp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rsp, "What do you mean, you called me!")
	})
	fmt.Println("Server started on 4000")
	http.ListenAndServe("localhost:4000", nil)
}
