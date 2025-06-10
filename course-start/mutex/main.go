package main

import (
	"sync"
)

var msg string
var mu sync.Mutex

func printstr(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	msg = str
	mu.Unlock()
}
func main() {
	// var wg sync.WaitGroup
	// msg = "hello world"

	// wg.Add(1)
	// go printstr("good morning", &wg)

	// go printstr("good evenig", &wg)
	// wg.Wait()
	// fmt.Println(msg)
	producerConsumer()
}
