package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	var users []User
	for i := 0; i < 10; i++ {
		users = append(users, User{i, "rashi", "rashi@gmail.com"})
	}
	w.Header().Set("Content-Type", "application/json") //not mandatory but good practice
	json.NewEncoder(w).Encode(users)
	fmt.Println("status of req is ", r.Method)

}

type userInfo struct {
	ID    int
	Name  string
	Email string
}

// fake db
var users = []userInfo{
	{ID: 1, Name: "Rashi", Email: "rashi@gmail.com"},
	{ID: 2, Name: "Akansha", Email: "akansha@gmail.com"},
	{ID: 3, Name: "akshu", Email: "akshu@gmail.com"},
	{ID: 4, Name: "Deepanshu", Email: "deepu@gmail.com"},
}

func getFromUrl(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	data := userInfo{
		Name:  name,
		Email: email,
	}
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(data)
	// fmt.Fprintf(nil, "get user data from url")

}

func upadteUser(w http.ResponseWriter, r *http.Request) {
	idParmas := chi.URLParam(r, "id")
	var updatedUser userInfo
	json.NewDecoder(r.Body).Decode(&updatedUser)

	for i, user := range users {
		if fmt.Sprint(user.ID) == idParmas {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email

			json.NewEncoder(w).Encode(users[i])

			fmt.Println("data", users[i])
			return
		}

	}

	http.Error(w, "error ", http.StatusNotFound)

}
func UserDeleted(w http.ResponseWriter, r *http.Request) {

	paramsId := chi.URLParam(r, "id")
	if len(users) > 1 {
		for i, user := range users {
			if fmt.Sprint(user.ID) == paramsId {
				users = append(users[:i], users[i+1:]...)
				fmt.Fprintf(w, "successfully deleted")
				fmt.Println("user delete ", users)
				return
			}
		}
	} else {
		fmt.Fprintf(w, "can not delete ")
	}

	http.Error(w, "error", http.StatusNotFound)
}

func main() {

	r := chi.NewRouter()

	r.Get("/getData", getUserData)
	r.Get("/get", getFromUrl)
	r.Put("/upadteUser/{id}", upadteUser)
	r.Delete("delete/{id}", UserDeleted)
	http.ListenAndServe(":8080", r)
}
