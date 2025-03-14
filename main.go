package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
}
