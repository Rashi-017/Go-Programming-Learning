package main

import "fmt"

var i, c, java bool

// var java int
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x int, y int) {
	x = sum / 2
	y = sum * 2
	return
}

func main() {
	java = true
	fmt.Println(add(3, 6))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split((40)))
	fmt.Println((java))

	var number = 1
	fmt.Println(number)
	fmt.Println("hello world")

	k := 3
	// short hand variable declaration is := this
	fmt.Println(k)
	var ToBe = false
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	var i int
	j := i
	fmt.Println((j))

	fmt.Printf("type %T and value id %v", j, j)
	fmt.Println()
	g := 0.867 + 0.5i
	fmt.Printf("type id %T and value id %v", g, g)
	fmt.Println()

}
