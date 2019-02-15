package main

import "fmt"

func main() {
	cards := newDeck()
	handCards, otherCards := deal(cards, 5)
	fmt.Println("Cards in hand")
	handCards.print()
	fmt.Println("Other Cards")
	otherCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
