package main

import (
	"fmt"
	"net/url"
)

const Url = "http://lco.dev:3000/learn?coursename=reactjs&paymentId=1234jfef"

func main() {
	fmt.Println("welcome to the web request ")
	result, _ := url.Parse(Url)
	fmt.Println(result)

}
