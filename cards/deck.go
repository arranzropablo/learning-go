package main

import "fmt"

type suite string

//We create a card type with a suite (which is a string) and a number (which is an int)
type card struct {
	number int
	suite suite
}

//Then we create the type deck which is a array of cards
type deck []card

//This is kind of a enum
var suites = [4]suite{"Spades", "Clubs", "Diamonds", "Hearts"}

//To avoid complexity we are creating a spanish deck which can be done with just numbers
func createDeck() deck {
	var deck deck
	for _, suite := range suites {
		for j := 1; j <= 12; j++{
			deck = append(deck, card{number: j, suite: suite})
		}
	}
	return deck
}

//This function has deck as a receiver
//This way any var of type deck in our application have access to func print
//This method is "attached" to deck type
//By convention names of receivers are just 1 or 2 letters
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}