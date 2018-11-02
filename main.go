package main

import ("fmt")

func main() {
	deck := newDeck()
	hand, remainingDeck := deal(deck, 5)
	hand.print()
	fmt.Println("")
	remainingDeck.print()
}
