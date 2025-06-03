package main

import "fmt"

func checkPalindrome(arr []int) bool {
	i := 0
	j := len(arr) - 1
	for i <= j {
		if arr[i] == arr[j] {
			i++
			j--
		} else {
			return false
		}

	}
	return true
}
func main() {

	fmt.Println("check the palindrome ")
	arr := []int{10, 30, 30, 40, 10}
	result := checkPalindrome(arr[:])
	fmt.Println(result)

}
