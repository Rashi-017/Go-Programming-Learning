package main

import (
	"fmt"
	"io"

	"strings"
)

func main() {
	// n, err := os.Open("./info.txt")
	n := strings.NewReader("helloworld")
	// if err != nil {
	// 	panic(err)
	// }

	total, err := countAlphabet(n)

	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}

func countAlphabet(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)
		for _, n := range buffer[:n] {
			if n >= 'a' && n <= 'z' || n >= 'A' && n <= 'Z' {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}

	}
}
