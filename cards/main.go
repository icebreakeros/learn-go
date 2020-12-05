/*
Package main is the main package of this program.
*/
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
