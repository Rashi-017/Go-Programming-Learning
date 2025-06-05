package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hellohandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "gest"
	}
	fmt.Fprintf(w, "hello,%s", name)

}
func createHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// var data map[string]string
	var data struct {
		Name  string
		Email string
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Created user: %+v", data)
	fmt.Println("data from post is ", data)

}

func main() {
	http.HandleFunc("/hello", hellohandle)
	http.HandleFunc("/create", createHandle)
	fmt.Println("make request on 'http://localhost:3000/hello'")
	http.ListenAndServe(":8000", nil)
}
