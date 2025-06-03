package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, val := range a {
		sum += val
	}
	c <- sum
}
func Channels() {
	var s = []int{10, 20, 30, 40, 50, 60}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("sum is ", x+y)

}
