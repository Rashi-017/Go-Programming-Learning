package main

import "fmt"

type Bot interface {
	getGreeting() string
}
type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}
	eb.getGreeting()
	sb.getGreeting()
	PrintGreeting(eb)
	PrintGreeting(sb)

}

func PrintGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
func (eb englishbot) getGreeting() string {
	return "Namestey Duniya "
}
func (sb spanishbot) getGreeting() string {
	return "Namestey spanish"
}
