package main

import (
	"fmt"
	"net/http"
)

type logwriter struct{}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("some err", err)
	}
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	//fmt.Println(string(bs))
	lw := logwriter{}
	lw.Write(bs)
}

func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
