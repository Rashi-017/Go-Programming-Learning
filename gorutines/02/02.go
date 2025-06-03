package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func processNum(numchan chan int) {
	for num := range numchan {
		fmt.Println("process number ", num)

		time.Sleep(time.Second)
	}

}

func sum(result chan int, n int, n2 int) {
	newResult := n + n2
	result <- newResult
}

func main() {
	numchan := make(chan int)

	go processNum(numchan)

	// for {
	// 	numchan <- rand.Intn(100) + 1
	// }

	result := make(chan int)
	go sum(result, 7, 5)

	res := <-result
	fmt.Println(res)

}
