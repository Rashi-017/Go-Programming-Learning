package main

import (
	"encoding/json"
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type user struct {
	Id    int
	Name  string
	Email string
}

var UserData []user

const filename = "data.txt"

func addUser(w http.ResponseWriter, r *http.Request) {

	var data user
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}
	filedata, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	var users []user
	if len(filedata) > 0 {
		err := json.Unmarshal(filedata, &users)
		if err != nil {
			http.Error(w, "Invalid data in file", http.StatusInternalServerError)
			return
		}
	} else {
		users = []user{}
	}
	for _, val := range users {
		if val.Id == data.Id {
			http.Error(w, "User ID already exists", http.StatusBadRequest)
			return
		}
	}

	UserData = append(users, data)
	newdata, err := json.MarshalIndent(UserData, "", " ")
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(filename, newdata, 0644)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(UserData)
	fmt.Println("User added:", data)

	//SaveToFile(filename, UserData)

	log.Println("successfull")
	fmt.Fprintf(w, "data added ")
	fmt.Println("data added is ", UserData)

}
func SaveToFile(filename string, userData []user) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer file.Close()
	data, err := json.MarshalIndent(userData, "", " ")
	if err != nil {
		panic(err)

	}
	file.Write(data)

}

func getAlUser(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "failed to read ", http.StatusInternalServerError)
		return
	}
	// newdata ,err  = json.Marshal(data)
	// json.NewEncoder(w).Encode(data)
	// fmt.Println("get data susccessfully", UserData)
	w.Write(data)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedData user
	json.NewDecoder(r.Body).Decode(&updatedData)

	id := chi.URLParam(r, "Id")

	filedata, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "file error", http.StatusInternalServerError)
		return
	}
	var users []user
	err = json.Unmarshal(filedata, &users)
	if err != nil {
		http.Error(w, "data err", http.StatusInternalServerError)
		return
	}
	updated := false

	for i, user := range users {
		if fmt.Sprint(user.Id) == id {
			users[i].Name = updatedData.Name
			users[i].Email = updatedData.Email

			updated = true
			break
			// json.NewEncoder(w).Encode(UserData[i])
			// fmt.Println("update data", UserData)
		}
	}
	if !updated {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	//save the updated data
	datatosave, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(filename, datatosave, 0644)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	// Step 4: Respond to client
	w.Header().Set("Content-Type", "application/json")
	w.Write(datatosave)
	fmt.Println("User updated successfully")

	// http.Error(w, "error", http.StatusNotFound)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "reading error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &UserData)

	if err != nil {
		http.Error(w, "invalid json", http.StatusInternalServerError)
	}

	paramsId := chi.URLParam(r, "id")
	log.Println("id recived ", paramsId)
	for i, user := range UserData {
		fmt.Println("Comparing with user ID:", user.Id)
		if fmt.Sprint(user.Id) == paramsId {
			UserData = append(UserData[:i], UserData[i+1:]...)
			SaveToFile(filename, UserData)
			fmt.Fprintf(w, "successfully deleted")
			fmt.Println("user delete ", UserData)

			return
		}

	}
	http.Error(w, "error", http.StatusNotFound)

}

func main() {

	router := chi.NewRouter()

	router.Post("/addUser", addUser)
	router.Get("/getall", getAlUser)
	router.Put("/update/{Id}", UpdateUser)
	router.Delete("/delete/{id}", DeleteUser)

	http.ListenAndServe(":8000", router)

}
