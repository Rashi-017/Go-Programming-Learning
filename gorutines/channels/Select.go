package main

import (
	"fmt"
	"time"
)

func Select() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "channel 1 is ready"
		close(ch1)

	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "channel 2 is ready "
		close(ch2)

	}()

	for i := 0; i < 10; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			fmt.Println("no channel is ready")

		}
		time.Sleep(500 * time.Millisecond)

	}
}
