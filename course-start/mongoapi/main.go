package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("server getting strated....")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("server started listening at port 3000")
}
