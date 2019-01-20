package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//To avoid complexity we are creating a spanish deck which can be done with just numbers
func newDeck() deck {
	var deck deck
	for _, suite := range suites {
		for j := 1; j <= 12; j++ {
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
		fmt.Println(card.String())
	}
}

//Receiver functions vs parameter functions:
//You cant declare more than 1 function with the same name if they have the same receiver (or no receiver)
//When using interfaces, the method executed is decided dinamically

func (d deck) deal(n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) save(name string) bool {
	error := ioutil.WriteFile(name, []byte(d.String()), 0666)
	if error == nil {
		return true
	}
	return false
}

func read(name string) deck {
	byted, err := ioutil.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		//return newDeck()
		os.Exit(1)
	}

	deckstring := string(byted)
	deckStr := strings.Split(deckstring, ",")

	var deck deck

	for _, c := range deckStr {
		deck = append(deck, toCard(c))
	}
	return deck
}

func toCard(s string) card {
	number, _ := strconv.Atoi(strings.Split(s, " of ")[0])
	suite := suite(strings.Split(s, " of ")[1])
	return card{number, suite}
}

//to restrict possivle suite values I could create a func "toSuite" and check the input

func (d deck) String() string {
	var s []string
	for _, c := range d {
		s = append(s, c.String())
	}
	return strings.Join(s, ",")
}

func (c card) String() string {
	return strconv.Itoa(c.number) + " of " + string(c.suite)
}

func (d deck) shuffle() {
	//If we don't generate a seed it will always be the same
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPos := r.Intn(len(d))

		d[newPos], d[i] = d[i], d[newPos]
	}
}
