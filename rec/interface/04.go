package main

import "fmt"

func acceptEveryThing(value interface{}) {
	fmt.Printf("type %T and value %d\n", value, value)
}
func function4() {
	acceptEveryThing("rashi")
	acceptEveryThing(true)
	acceptEveryThing([]int{1, 2, 4})

}
