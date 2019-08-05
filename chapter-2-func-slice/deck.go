package main

import "fmt"

type deckType []string

func newDeck () deckType {
	cards := deckType{}

	cardSuits := []string{ "Spade", "Club", "Diamond", "Heart" }
	cardValues := []string{ "Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King" }

	for _, values := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, values +" of "+suit)
		}
	}

	return cards
}

// Receiver function
func (d deckType) print () {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// Multiple return
func (d deckType) deal (handSize int) (deckType, deckType) {
	return d[:handSize], d[handSize:]
}
