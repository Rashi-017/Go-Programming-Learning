package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("server start")
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		fmt.Println(name)
		w.Write([]byte("Hello, " + name + "!"))
	})

	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Name  string
			Email string
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "data is %v", data)
		fmt.Println("data is ", data)

	})

	http.ListenAndServe(":3000", r)

}
