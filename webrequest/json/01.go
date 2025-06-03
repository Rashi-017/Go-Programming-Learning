package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string //`json:"name"`
	Age    int    //`json:"age"`
	Active bool   //`json:"active"`
}

func main() {
	fmt.Println("learn json data")
	p := Person{"john", 23, true}
	jsonres, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonres))
	var person Person
	json.Unmarshal(jsonres, &person)
	fmt.Println(person)
}
