package main

import "fmt"

var i interface{}

func found(i interface{}) {
	fmt.Printf("value %d and type is %T\n", i, i)
}

func function3() {
	s := "simlelearn"
	found(s)
}
