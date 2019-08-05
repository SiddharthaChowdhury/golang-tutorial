package main

func main() {
	deck := newDeck()

	hand, _ := deck.deal(5)
	hand.print()
}
