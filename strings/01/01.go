package main

import "fmt"

func countPalindrome(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			count++
		}
	}
	return count
}

func main() {

	fmt.Println("count the vowels in string")
	var str = "aeiou"
	result := countPalindrome(str)
	fmt.Println(result)

}
