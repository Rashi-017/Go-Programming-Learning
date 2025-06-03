package main

import (
	"fmt"
)

type ageError struct {
	age int
}

func (e *ageError) Error() string {
	return fmt.Sprintf("Too young to drive: age %d", e.age)
}

func CheckDriveOrNot(age int) (bool, error) {
	if age <= 18 {
		return false, &ageError{age: age}
	} else {
		return true, nil
	}

}

func DriveOrNot() {
	var age int
	fmt.Println("enter your  age")
	fmt.Scan(&age)
	_, err := CheckDriveOrNot(age)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("you can drive your age  is greater than 18 ")

}
