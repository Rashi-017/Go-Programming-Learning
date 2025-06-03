package main

import (
	"fmt"
)

func findMax(arr []int) int {
	var great = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] >= great {
			great = arr[i]
		}
	}
	return great
}

func main() {
	fmt.Println("enter the values of array ")
	var val int

	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&val)
		arr[i] = val
	}

	fmt.Println(arr)
	result := findMax(arr[:])
	fmt.Println(result)

}
