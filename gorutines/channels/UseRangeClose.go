package main

import "fmt"

func UseRangeClose() {
	fmt.Println("use range and close in channel")
	c := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}

	ch := make(chan string, 10)
	fmt.Println("enter the value ")
	var input string

	for i := 0; i < 5; i++ {
		fmt.Scan(&input)
		ch <- input
	}
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}

	val, ok := <-ch
	if !ok {
		fmt.Println("channel is closed")
	}
	fmt.Println(val)

}
