package main

import (
	"fmt"
)

func main() {
	d := newDeck()
	//d.print()

	d.shuffle()

	playerdeck, d := d.deal(5)

	//When calling a func with a pointer as receiver we can call it with the original var (without &) and it will do the workaround
	//playerdeck := (&d).deal(5)
	//playerdeck := d.deal(5)

	fmt.Println("Player 1 deck is:")
	playerdeck.print()

	//&<var> is the memory address of a variable (a pointer)
	//*<var> is the value of a memory address
	//*<type> is a pointer which points a that type

	//playerdeck.save("playerdeck.txt")
	//fmt.Println("Saved file:")
	//read("playerdeck.txt").print()

}
