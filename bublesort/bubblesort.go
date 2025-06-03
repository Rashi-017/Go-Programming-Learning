package main

import "fmt"

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {

	fmt.Println("bubble sort ")
	// var arr [5]int
	// var val int
	// fmt.Println("enter the arr value")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Scan(val)
	// 	arr[i] = val
	// }
	arr := []int{50, 40, 30, 20, 10}

	result := bubblesort(arr[:])
	defer fmt.Println(result)

	type v struct {
		Name string
		Age  int
	}
	vertex := v{"rashi", 23}

	p := &vertex
	fmt.Println(p.Name)

}
