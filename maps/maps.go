package main

import "fmt"

func main() {

	fmt.Println("start with map")

	language := make(map[string]string)
	language["py"] = "python"
	language["js"] = "javascript"
	language["rb"] = "ruby"
	language["rb"] = "rails"

	fmt.Println(language)
	//delete(language, "rb")
	fmt.Println(language)

	// for key, value := range language {
	// 	fmt.Printf("for key %v and value is %v", key, value)
	// }
	// var countlogin = 23
	// var result string
	// if countlogin < 10 {
	// 	result = "regular user"
	// } else if countlogin > 10 {
	// 	result = "watch out user"
	// } else {
	// 	result = "no user"
	// }

	// fmt.Println(result)

	if num := 3; num < 5 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	hitesh := User{"rashi", "rashigmail.com", 23}

	fmt.Println(hitesh.Name)
	fmt.Printf("type of hitesh is %T", hitesh)

	//IF ELSE
	var countlogin = 23
	var result string
	if countlogin < 10 {
		result = "regular user"
	} else if countlogin > 10 {
		result = "watch out user"
	} else {
		result = "no user"
	}

	fmt.Println(result)

	var val = make(map[string]vertex)
	val["hello"] = vertex{23, 45}
	fmt.Println(val)

	var value = map[string]vertex{
		"good":  vertex{34, 50},
		"hello": vertex{32, 90},
	}
	fmt.Println(value)

	v, k := val["yes"]
	fmt.Printf("key value %d exist %t\n", v, k)

}

type User struct {
	Name       string
	Email      string
	Rollnumber int
}

type vertex struct {
	ltd  int
	long int
}
