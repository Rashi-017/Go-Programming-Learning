package main

import (
	"fmt"
	"time"
)

func main() {

	// while loop in go is like :

	// for sum < 10 {
	// 	fmt.Print(sum)
	// 	sum += 1
	// }

	if sum := 30; sum > 20 {
		fmt.Println("hello world")
	}

	switch color := "blue"; color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "yellow":
		fmt.Println("color is yellow")
	default:
		fmt.Println("please try again ")
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 4:
		fmt.Println("saturday.")
	default:
		fmt.Println("Too far away.")
	}
	//without any condition switch is true

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good aftoornoon")
	default:
		fmt.Println("no wish")
	}

	//defer
	fmt.Println("hello")

	fmt.Println("world")

	//defer

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)

	// }
	// fmt.Println("hello")
	var flower string = "sunflower"
	switch flower {
	case "sunflow":
		fmt.Println("sunflower")
	default:
		fmt.Println("no flower found")

	}
	function2()
	function3()

}
