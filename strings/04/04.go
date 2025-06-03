package main

import "fmt"

func targetsum(arr []int, target int) (x int, y int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return arr[i], arr[j]
			}
		}
	}
	return -1, -1
}
func countpair(arr []int, target int) int {
	pair := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pair++
			}
		}
	}
	return pair
}
func main() {

	fmt.Println("target sum problem ")
	arr := []int{3, 9, 8, 2, 4, 2, 1}
	target := 10
	x, y := targetsum(arr[:], target)
	fmt.Printf("pair found is %d and %d\n", x, y)

	result := countpair(arr[:], 10)
	fmt.Println(result)

}
