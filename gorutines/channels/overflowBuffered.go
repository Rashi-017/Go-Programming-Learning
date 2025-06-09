package main

import (
	"fmt"
	"time"
)

func OverflowBuffered() {

	ch := make(chan string, 2)
	ch <- "string1"
	ch <- "string2"

	ch1 := make(chan int, 2)

	// ch <- "string3"
	go func() {
		for i := 0; i < 2; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	for val := range ch1 {
		fmt.Println(val)
	}
	select {
	case ch <- "string1":
		fmt.Println("string 1")
	case ch <- "string2":
		fmt.Println("string 2")
	case ch <- "string3":
		fmt.Println("string3")
	default:
		time.Sleep(time.Second * 2)
		fmt.Println(" channel is full ")

	}

	fmt.Println("this line will never reach ")
}
