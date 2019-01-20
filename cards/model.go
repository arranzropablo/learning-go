package main

type suite string

//This is kind of a enum (Other way: https://programming.guide/go/define-enumeration-string.html)
//Would be good to find a way to restrict values of type suite to these 4, for example
var suites = [4]suite{"Spades", "Clubs", "Diamonds", "Hearts"}

//We create a card type with a suite (which is a string) and a number (which is an int)
type card struct {
	number int
	suite  suite
}

//Then we create the type deck which is a array of cards
type deck []card

//Enum can be also done this way (should add iota and more stuff)
/*
//This is kind of a enum
const (
	SPADES suite = "Spades"
	CLUBS suite = "Clubs"
	DIAMONDS suite = "Diamonds"
	HEARTS suite = "Hearts"
)

var suites = [4]suite{SPADES, CLUBS, DIAMONDS, HEARTS}
*/
