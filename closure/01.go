package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {

	fmt.Println("what are clousers")
	result := adder()
	ans := result(3)
	fmt.Println(ans)
}
