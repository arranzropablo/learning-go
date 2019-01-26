//only files under package main compiles to executable, others under other package can be use as helpers though
package main

import (
	//We cant declare variables or import packages which wont be used, because it wont compile
	"fmt"
)

func main() {
	//fmt is the standard i/o system
	//func can return more than one param, for example Println returns num of chars printed and error
	fmt.Println(fmt.Println("Hello world"))

	//Sprintf formats the given string and returns it or an error
	fmt.Println(fmt.Sprintf("Hello world %d", 5))

	//Types and vars

	//type is optional (if it has initial value, if not it will be initialized to "zero value" i.e. empty string, false...)
	var a string = "This is a string"
	var (
		text    string = "You can define more than"
		number  int    = 1
		text2   string = "variable in this block"
		boolean bool   = true
	)
	var asd, dsa string = "this is", "possible too"

	//short variables are a short way to initialize variables
	short := "short"
	shortb := true

	//you can initialize many variables of dif types at once
	boolvar, intvar, floatvar := true, 64, 64.2

	fmt.Println(a, text, number, text2, asd, dsa, boolean, short, shortb, boolvar, intvar, floatvar)

	const CONSTANT = "GOPHER"

	fmt.Println(getAOfClubs())
	deckSize := deckSize()
	fmt.Println(deckSize)

	suites := []string{"Hearts", "Clubs", "Spades"}

	//append, delete, etc... does not modify the object, but return a copy (we can reassign then)
	suites = append(suites, "Diamonds")
	size := len(suites)
	fmt.Println(suites, size)

	//range returns the index and the object
	//if we just want the object we will get an error for not using the index var
	//we can solve it "naming" the index with an underscore
	for _, suite := range suites {
		fmt.Println(suite)
	}
}

//functions can return stuff if return types are specified between parentheses (if more than one)
func getAOfClubs() (string, string) {
	return "Ace", "Clubs"
}

func deckSize() int64 {
	return 52
}

const EXAMPLE = "You can declare const and vars outside of functions, if they are not short variables"
