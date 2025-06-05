package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
}

func channelPractice() {
	// ch := make(chan int, 2)
	// close(ch)
	// _, ok := <-ch
	// if !ok {
	// 	fmt.Println("Channel closed")
	// }
	// ch <- 1
	// fmt.Println(<-ch)

	// p:=person{"rashi"}
	// stop := make(chan struct{})

	// go func() {
	// 	select {
	// 	case <-stop:
	// 		fmt.Println("Stopping")
	// 	}
	// }()

	// close(stop) // broadcast stop to all listening goroutines
	ch := make(chan int)
	// here code is block
	// ch <- 9

	go func() {
		ch <- 9
	}()
	select {
	case msg := <-ch:
		fmt.Println("recived", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("no value recived timeout")

	}

}
