package main

import (
	"fmt"
	"time"
)

func task(n int) {
	fmt.Println("task", n)
}

func main() {
	// for i := 0; i <= 10; i++ {
	// 	go task(i)
	// }

	for i := 0; i <= 10; i++ {
		//anonymus function
		func() {
			go fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 2)

}
