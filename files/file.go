package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"os"
)

func main() {

	fmt.Println("hello")
	content := "hello my name is rashi how are you "

	file, err := os.Create("./dummy.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of file is ", length)

	defer file.Close()
	readfile("./dummy.txt")

}
func readfile(filepath string) {
	databyte, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("data from file is ", string(databyte))
}
