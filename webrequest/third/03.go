package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("make get request")

	res, err := http.Get(Url)
	if err != nil {
		fmt.Println("some thing went wrong", err)
	}

	resdata, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("someerror ", err)
	}

	// fmt.Println(string(resdata))

	result, err := url.Parse(Url)
	if err != nil {
		panic(err)
	}

	result.RawQuery = "userId=10"
	newUrl := result.String()

	fmt.Println("new url", newUrl)
	res, err = http.Get(newUrl)
	resdata, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(string(resdata))

	

}
