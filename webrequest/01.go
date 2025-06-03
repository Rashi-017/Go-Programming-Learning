package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Url = "https://dummy-json.mock.beeceptor.com/todos"

func main() {

	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	//data := response.Body
	result, err := url.Parse(Url)
	if err != nil {
		panic(err)
	}
	fmt.Println("result", result.Scheme)
	fmt.Println("host ", result.Host)
	fmt.Println("query in url", result.RawQuery)
	
	fmt.Println("response is ", response.Status)
	databyte, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()
	content := string(databyte)
	fmt.Println(content)

}
