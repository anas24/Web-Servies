package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hears"}

	cardValues := []string{"Ace", "two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "ten", "King", "Joke", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " Of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
