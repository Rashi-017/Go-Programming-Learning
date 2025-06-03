package main

import (
	"fmt"
	"strings"
	"sync"
)

type wordCounter struct {
	count map[string]int
}

// Increment safely increments the word count
func (w *wordCounter) Increment(word string) {

	w.count[word]++

}

// Value safely gets the count of a word
func (w *wordCounter) value(word string) int {
	return w.count[word]
}

func mutex3() {
	fmt.Println("Map with Mutex")
	text := "go is simple but powerful and go is fast"
	words := strings.Fields(text)
	counter := wordCounter{count: make(map[string]int)}

	var wg sync.WaitGroup

	for _, val := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			counter.Increment(w)
		}(val)
	}
	wg.Wait()

	fmt.Println("Word counts:")
	for word, count := range counter.count {
		fmt.Printf("%s: %d\n", word, count)
	}

}
