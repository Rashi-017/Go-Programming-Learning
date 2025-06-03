package main

import (
	"fmt"
)

type small struct {
	x int
	y int
}
type value struct {
	name string
	age  int
}

func function2() {

	v := small{1, 2}
	fmt.Println(v.x)

	i := value{"rashi", 23}
	p := &i
	fmt.Println(p.name)

	// i, j := 5, 6

	// p := &i
	// q := &j

	// fmt.Println(*p)
	// fmt.Println(*q)
	// *p = 21
	// fmt.Println(i)

}
