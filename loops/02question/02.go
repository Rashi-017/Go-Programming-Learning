package main

import "fmt"

func reverseslice(arr []int) [5]int {
	i := 0
	j := len(arr) - 1
	for i <= j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
	return [5]int(arr)

}

func main() {

	fmt.Println("reverse the slice ")
	fmt.Println("enter the input ")
	var val int
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&val)
		arr[i] = val
	}
	fmt.Println(arr)

	result := reverseslice(arr[:])
	fmt.Println(result)

}
