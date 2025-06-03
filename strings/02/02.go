package main

import "fmt"

func reverseTheString(r []rune) string {

	// if len(r) > 1 {
	// 	r[0], r[1] = r[1], r[0]
	// }
	i := 0
	j := len(r) - 1
	for i <= j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
func main() {
	fmt.Println("reverse the string")
	str := "rashi"
	fmt.Println(str)
	runes := []rune(str)
	result := reverseTheString(runes)

	fmt.Println(result)

}
