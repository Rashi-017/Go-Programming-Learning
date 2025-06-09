package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/hello", getresp)
	ser := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}
	//http.ListenAndServe(":3000",r)
	ser.ListenAndServe()
}
func getresp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world how are you </h1><img height='500px'width='500px' src='https://pressbooks.cuny.edu/app/uploads/sites/93/2022/08/thanuj-mathew-8CSTVoDMEXg-unsplash-scaled.jpg'> </img><button onclick=clickfunc>click me </button>"))
}
func clickfunc() string {
	return "hello world"
}
