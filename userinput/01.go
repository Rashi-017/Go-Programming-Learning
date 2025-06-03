package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	"strings"
)

func main() {

	// fmt.Println("give me the rating of pizza")
	// reader := bufio.NewReader(os.Stdin)

	// //comma ok
	// input, _ := reader.ReadString('\n')

	// fmt.Println("thanks for rating ", input)

	// newNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("one add to your rating ", newNumber+1)
	// }

	// num := 23
	// var ptr = &num
	// *ptr = num + 2
	// fmt.Println(num)

	// var newarr [4]string

	// for i := 0; i < 4; i++ {
	// 	input, _ := reader.ReadString('\n')
	// 	newarr[i] = input
	// }

	// fmt.Println(newarr)
	// var slice1 = newarr[:2]
	// slice1[0] = "pineapple"
	// //var index int = 2
	// //newarr = append(newarr[:2], newarr[index+1:]...)
	// fmt.Println(newarr)

	//slice of slices

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][1] = "x"
	board[1][0] = "0"
	board[2][2] = "x"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
