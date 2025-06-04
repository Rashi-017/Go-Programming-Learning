package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
	// "net/url"
)

const Url = "https://jsonplaceholder.typicode.com/posts"

type Todo struct {
	UserId int
	Id     int
	title  string
	Body   string
}

func PerformGetRequest() {
	fmt.Println("posts data")
	res, err := http.Get(Url)
	if err != nil {
		fmt.Println("404 not found")
	}
	// resdata, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("something went wrong", err)
	// }
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("something went wrong")
		return
	}
	fmt.Println(todo)

}

type UserData struct {
	UserId   int    `json:"UserId"`
	Title    string `json:"Title"`
	Compelte bool   `json:"Complete"`
}

func PerformPostRequest() {
	UserData := UserData{
		UserId:   1234,
		Title:    "ramesh",
		Compelte: true,
	}
	jsondata, err := json.Marshal(UserData)
	if err != nil {
		fmt.Println("error", err)
	}
	jsonstring := string(jsondata)
	jsonReader := strings.NewReader(jsonstring)
	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", jsonReader)
	if err != nil {
		fmt.Println("error ", err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	fmt.Println(res.Status)
}
func main() {
	//PerformGetRequest()
	PerformPostRequest()

}
