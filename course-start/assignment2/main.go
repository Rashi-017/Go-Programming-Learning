package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	f, err := os.Open(os.Args[0])
	if err != nil {
		fmt.Println("error", err)
	}
	io.Copy(os.Stdout, f)

}
