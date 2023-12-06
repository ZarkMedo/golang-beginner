// json demo

package main

import (
	"encoding/json"
	"fmt"
)

// struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	HaveSon bool `json:"have_son,omitempty"` // if the value is empty, omit it
}

func main() {
	// create struct
	p := Person{
		Name: "John",
		Age:  30,
	}

	// marshal to json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	p2 := Person{
		Name: "Johnnnnnn",
	}
	// unmarshal from json
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
