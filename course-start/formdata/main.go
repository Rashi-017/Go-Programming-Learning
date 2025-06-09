package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("form data")
	data := url.Values{}
	data.Add("firstame", "rashi")
	data.Add("lastname", "sharma")
	data.Add("age", "23")

}
