package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const Url = "https://jsonplaceholder.typicode.com/posts"

func getreq() {
	fmt.Println("get call ")

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("error", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(data))
}

type Postsdata struct {
	Id        int
	userId    int
	Title     string
	completed bool
}

func postreq() {
	todo := Postsdata{
		Id:        23,
		userId:    490,
		Title:     "new Todo",
		completed: true,
	}
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("some err", err)
	}
	jsonReader := strings.NewReader(string(jsondata))
	res, err := http.Post(Url, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(res.Status)

}

func updatereq() {
	todo := Postsdata{
		Id:        23,
		userId:    490,
		Title:     "new Todo",
		completed: true,
	}
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("err", err)
	}

	jsonReader := strings.NewReader(string(jsondata))
	const myurl = "https://jsonplaceholder.typicode.com/posts/1"
	//update request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("some error", err)
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

}
func deletereq() {
	const myurl = "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("some err", err)
		return
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)

	}
	data, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(data))
	defer res.Body.Close()
	fmt.Println(res.Status)

}

func main() {
	getreq()
	postreq()
	deletereq()
	updatereq()

}
