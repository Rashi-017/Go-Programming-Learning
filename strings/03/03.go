package main

import (
	"fmt"
)

func countfreq(str string) map[string]int {
	runes := []rune(str)
	freqmap := make(map[string]int)

	for _, val := range runes {
		ch := string(val)
		freqmap[ch]++
	}
	return freqmap

}

func main() {

	fmt.Println("count the frequency of each word in string ")
	str := "racecar"
	fmt.Println(str)
	result := countfreq(str)
	fmt.Println(result)
	// fmt.Printf("%d ----%d\n",result)

}
