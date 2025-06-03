package main

import (
	"fmt"

	"time"
)

func Select2() {
	// ch1:= make(chan int)

	ch2 := make(chan int)

	// go func ()  {
	// 	time.Sleep(time.Second*1)
	// 	ch1<-20
	// }()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 40
	}()

	select {

	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("no channel is ready  yet ")

	}
	time.Sleep(time.Second)

}
