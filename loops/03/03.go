package main

import "fmt"

func checkduplicate(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if val == arr[j] {
				return true
			}
		}
	}
	return false
}

// create freq map
func createfreqMap(arr []int) map[int]int {
	freq := make(map[int]int)

	for _, val := range arr {
		freq[val]++
	}
	return freq
}
func main() {
	arr := []int{20, 30, 10, 10, 10}
	result := checkduplicate(arr[:])
	fmt.Println(result)

	resultMap := createfreqMap(arr[:])
	fmt.Println(resultMap)
	for key, val := range resultMap {
		fmt.Printf("%d --%d\n", key, val)
	}

	// key := make([]int, 0, len(resultMap))
	// for k := range resultMap {
	// 	key = append(key, k)

	// }
	// for _, k := range key {
	// 	fmt.Printf("%-10d %-10d\n", k, resultMap[k])
	// }

}
