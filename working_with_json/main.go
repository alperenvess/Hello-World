package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName   string `json:"name"`
	LastName    string
	Internal    string `json:"-"`
	Mandatory   int    `json:"mandatory"`
	Zero        int    `json:"zero,omitempty"`
	iDoNotSeeIt int    `json: "notSeen"`
}

func main() {
	input := `{
		"name": "Milena",
		"lastName": "Rogowska"
	}`

	var empl Employee

	err := json.Unmarshal([]byte(input), &empl)
	if err != nil {

		panic(err)
	}

	fmt.Println(empl.FirstName)
	fmt.Println(empl.LastName)
	fmt.Println("%v", empl)

	empl.Mandatory = 0
	empl.Zero = 0

	out, err := json.Marshal(empl)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(out))

}
