package main

import "fmt"

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("processing.....")
}

func send(email chan string) {

}

func main() {

	done := make(chan bool)
	go task(done)

	<-done

}
