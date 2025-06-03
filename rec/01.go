package main

import "fmt"

func fib(x int) int {
	if x <= 1 {

		return x
	}
	fmt.Print(x)
	return fib(x-1) + fib(x-2)
}

type Person struct {
	name     string
	age      int
	location string
}

// pointer reciver
func (p *Person) set_val() {
	fmt.Scan(&p.name)
	fmt.Scan(&p.age)
	fmt.Scan(&p.location)
}

// value reciver
func (p Person) get_val() {
	fmt.Println("name:", p.name)
	fmt.Println("age :", p.age)
	fmt.Println("location :", p.location)

}

// func (v vertex) abs()float64{
// 	return (v.x * v.y)
// }

func main() {

	fmt.Println("fibonacci")
	fmt.Println(fib(5))
	per := Person{}
	per.set_val()
	per.get_val()
}
