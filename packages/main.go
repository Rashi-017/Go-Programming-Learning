package main

import (
	"fmt"
	"packages/auth"
)

func main() {
	auth.LoginUser()
	auth.Register()
	newUser := auth.User{
		Id:      12390,
		Name:    "rashi",
		Roll_no: 017,
		Email:   "@gmail.com",
	}

	fmt.Println("newUser created is ", newUser)
}
