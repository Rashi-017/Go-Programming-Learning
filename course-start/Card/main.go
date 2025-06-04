package main

// import "fmt"

import "fmt"

func main() {
	cards := newdeck()

	fmt.Println(cards.tostring())
	cards.print()
	hand, remainingDeck := Deal(cards, 3)

	hand.print()
	remainingDeck.print()
	cards.SaveToFile("my_card")
	readfile("my_card")
	cards.shuffel()
	cards.print()

}
