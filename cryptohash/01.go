package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "this is my password"

	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println("before :", s)
	fmt.Printf("after : %x\n", bs)
}
