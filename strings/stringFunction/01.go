package main

import (
	"fmt"
	"strings"
	// s "strings"
)

func main() {

	fmt.Println("sring functions ")
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("count", strings.Count("test", "t"))
	fmt.Println("hasPrefix", strings.HasPrefix("test", "te"))
	fmt.Println("hasSuffix", strings.HasSuffix("rashi", "hi"))
	fmt.Println("Index", strings.Index("rashi", "h"))
	fmt.Println("join", strings.Join([]string{"Rashi", "sharma"}, "/"))
	fmt.Println("repeat", strings.Repeat("r", 3))
	fmt.Println("replace", strings.Replace("rashi", "h", "t", -1))
	fmt.Println("replace", strings.Replace("foo", "o", "0", 1))
	fmt.Println("split", strings.Split("03/june/2004", "/"))
	fmt.Println("tolowercase", strings.ToLower("Rashi"))
	fmt.Println("toupprecase", strings.ToUpper("rashi"))

}
