package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	ch := make(chan string)
	for _, val := range links {

		go checkStatus(val, ch)
		time.Sleep(time.Second * 2)

	}
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// for {
	// 	go checkStatus(<-ch, ch)
	// }
	// for l:=range ch{
	// 	go checkStatus(l,ch)
	// }
	// for l := range ch {
	// 	go func() {
	// 		time.Sleep(time.Second)
	// 		checkStatus(l, ch)
	// 	}()
	// }
	channelPractice()

}
func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down ")
		c <- link
		return
	}
	fmt.Println(link, "link is up")
	c <- link
}
