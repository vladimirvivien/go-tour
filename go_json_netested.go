package main

import (
	"encoding/json"
	"fmt"
)

type FileStatus struct {
	Size int
}

type FileStatuses struct {
	FileStatus []FileStatus
}

type Statuses struct {
	FileStatuses FileStatuses
}

const data =`
{
	"FileStatuses":{
		"FileStatus":[
			{"Size":100}
		]
	}
}
`

func main() {
	var statuses Statuses
	json.Unmarshal([]byte(data), &statuses)
	fmt.Printf("%v", statuses)
}