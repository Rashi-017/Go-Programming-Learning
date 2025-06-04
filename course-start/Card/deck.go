package main

import (
	"fmt"
	"io/ioutil"

	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, val := range d {
		fmt.Printf("%d -- %s\n", i, val)
	}
}

func newdeck() deck {
	cards := deck{}

	cardSuit := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValue := []string{"one", "two", "three", "four"}
	for _, suit := range cardSuit {
		for _, val := range cardValue {
			cards = append(cards, val+" of "+suit)
		}

	}
	return cards

}

func Deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]

}
func (d deck) tostring() string {
	return strings.Join([]string(d), ",")

}
func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.tostring()), 0666)
}

func readfile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(0)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}
func (d deck) shuffel() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
