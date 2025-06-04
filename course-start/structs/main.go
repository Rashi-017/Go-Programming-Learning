package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
}

func main() {
	Alex := Person{firstname: "Alex", lastname: "sharma"}
	fmt.Println(Alex)
	Alex.firstname = "Rashi"
	fmt.Println(Alex)
	fmt.Println(Filldetails("rashi", "sharma", "rashi@gmail.com", 456796987))
	Alex.changeInfo("sushil")
	fmt.Println(Alex)
}
