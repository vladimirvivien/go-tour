package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Version int
	Type    int
	Body    string
}

func main() {
	// unmarshalling example
	data := []byte(`{"Version":1234455, "Type":1, "Body":"Hello World!"}`)
	var msg Message
	json.Unmarshal(data, &msg)
	fmt.Println(msg)

	data2, _ := json.Marshal(Message{Version: 98763, Type: 0, Body: "!World Hello"})
	fmt.Println(data2)

	// generic decoding
	str := `{"Version":45, "Type":1, "Body":"Hello Universe."}`
	var data3 map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data3); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("JSON as generic map", data3)
	fmt.Printf("map[%s] = %v \n", "Version", data3["Version"].(float64))
	fmt.Printf("map[%s] = %v \n", "Body", data3["Body"].(string))
}
