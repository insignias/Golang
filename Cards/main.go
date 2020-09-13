package main

func main() {
	//cards := newDeck()

	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()
	//cards.saveToFile("mycards")

	cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()
}
