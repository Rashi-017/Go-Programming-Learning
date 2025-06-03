package main

import (
	"fmt"
)

var f interface{} = "sushil"

func describe(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case bool:
		fmt.Printf("%q is bytes long\n", v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Println("no datatype match")
	}

}

func function5() {
	val, ok := f.(int)
	fmt.Println(val, ok)

	describe(true)

}
