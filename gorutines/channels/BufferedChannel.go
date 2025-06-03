package main

import (
	"fmt"
	"time"
)

func BufferedChannel() {
	//now buffered have fixed capacity

	ch := make(chan string, 2)

	fmt.Println("worker send the box 1")
	ch <- "Box-1"
	fmt.Println("worker send the box 2 ")
	ch <- "Box-2"
	fmt.Println("we can not add box 3 in it ")
	ch <- "box3"

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Box 1 is received ", <-ch)
		fmt.Println("box 2 is recived ", <-ch)
		fmt.Println("box 3 is recived ", <-ch)

	}()

	time.Sleep(time.Second * 3)
}
