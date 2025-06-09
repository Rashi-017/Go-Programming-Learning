package main

import (
	"fmt"
)

	type voice interface {
		say() string
	}
type dog struct{}

func (c dog) say() string {
	return "dog voice bark"
}

type cat struct{}

func (c cat) say() string {
	return "cat voice is meeao"
}

//	func printArea(s voice) {
//		fmt.Println(s.say())
//	}
func main() {
	c := cat{}
	d := dog{}
	// printArea(c)
	// printArea(d)
	c.say()
	d.say()
	fmt.Print("interface")
	function2()
	function3()
	function4()
	function5()
}
