package main

import (
	"fmt"
)

func main() {
	d := newDeck()
	//d.print()

	d.shuffle()

	playerdeck, d := d.deal(5)

	fmt.Println("Player 1 deck is:")
	playerdeck.print()

	//playerdeck.save("playerdeck.txt")
	//fmt.Println("Saved file:")
	//read("playerdeck.txt").print()

}
