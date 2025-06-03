package main

import "fmt"

func UnbufferedChannel() {
	ch := make(chan string)

	//packer
	go func() {
		box := <-ch
		fmt.Println("Packer recive the box", box)
	}()

	//worker
	ch <- "Box-1"
	fmt.Println("worker  send the box ")

}
