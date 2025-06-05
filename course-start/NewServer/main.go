package main

import (
	"fmt"
	"log"
	"os"
    "github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")
	godotenv.Load()
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("poort is not found in env fle")
	}
	fmt.Println("port", portstring)

	router := chi.NewRouter()

	srv:= &http.Server{
		handle:router,
		Addr:":"+portstring
	}
	err:= srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}
