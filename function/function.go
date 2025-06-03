package main

import "fmt"

func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hello pagal"
}

// func greet(name string, wish string) string  {
// 	 return ("hello world %d and %d",name,wish)
// }

func print[T any](arr []T) {
	fmt.Println(arr)
}

func main() {

	fmt.Println("start with func")
	fmt.Println(proadder(2, 3, 4, 5, 5))
	_, value2 := proadder(2, 4, 5, 7)
	// fmt.Println(value1)
	fmt.Println(value2)
	//fmt.Println(greet("rashi", "goodmorning"))

	hitesh := User{"rashi", 23, true}

	hitesh.changevalue()

	var val = []int{10, 20, 30, 40}
	print(val)

}

type User struct {
	Name   string
	Age    int
	Status bool
}

func (u User) changevalue() {
	// u.name = "sakshi"
	fmt.Println("hello ", u.Name)

}
