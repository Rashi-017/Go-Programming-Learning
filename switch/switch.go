package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//var num = rand.Intn(6) + 1

	switch num := rand.Intn(4); num {
	case 1:
		fmt.Printf("dice value  %d is open you can open", num)
	case 2:
		fmt.Println("you can write ")
	case 3:
		fmt.Println("you can watch the file ")
	default:
		fmt.Println("thank you come again")
	}

	var days = []string{"sunday", "monday", "tuesday", "wednesday"}
	for index, val := range days {
		fmt.Printf("index %v and value is %v\n", index, val)
	}
	for _, day := range days {
		fmt.Printf("days are %v\n", day)
	}

	//continue and break

	for a := 1; a < 10; a++ {
		if a == 2 {
			goto loco
		}
		if a == 5 {
			// continue
			continue
		}
		fmt.Println(a)
		a++
	}
loco:
	fmt.Println("hello world ")

}
