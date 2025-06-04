package main

import (
	"fmt"
)

func main() {
	fmt.Println("create a map")
	colors := map[string]string{
		"red":    "ff000",
		"gren":   "4bf745",
		"blue":   "fh409",
		"yellow": "jio8",
	}
	fmt.Println(colors)
	//delete(colors, "gren")
	//fmt.Println(colors)
	for key, val := range colors {
		fmt.Printf("at %v value is %v\n", key, val)
	}
}
