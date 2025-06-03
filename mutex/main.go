package main

import (
	"fmt"
	"sync"
)

type posts struct {
	view int
	mu   sync.Mutex
}

func (v *posts) countviews(wg *sync.WaitGroup) {

	// defer func() {

	wg.Done()
	// }()
	v.mu.Lock()
	v.view += 1
	v.mu.Unlock()
}

func main() {
	fmt.Println("mutex the resource ")
	mypost := posts{view: 0}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.countviews(&wg)

	}

	wg.Wait()
	fmt.Println("total is ", mypost.view)
	mutex()
	mutex3()
}
