//only files under package main compiles to executable, others under other package can be use as helpers though
package main

import (
	//We cant declare variables or import packages which wont be used, because it wont compile
	"fmt"

	//We can import libraries from github with the route and with "go get <URL>"
	//"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
)

const EXAMPLE = "You can declare const and vars outside of functions, if they are not short variables"

//We can give constants a number value with iota (starting in 0 and adding 1 each const), this is similar to enums
const (
	a = iota // 0
	b        // 1
	c        // 2
)

//With iota we can skip an iota
const (
	d = iota // 0
	_        // 1
	f        // 2
)

//<< and >> are binary shift
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	//You can defer a func, it means it will be executed when the surrounding func returns
	defer fmt.Println("defeeeeeeeeeeeeer")

	//fmt is the standard i/o system
	//func can return more than one param, for example Println returns num of chars printed and error
	fmt.Println(fmt.Println("Hello world"))

	//Sprintf formats the given string and returns it or an error
	fmt.Println(fmt.Sprintf("Hello world %d", 5))

	//We can format text with multiple utilities, like converting a number to binary or hex
	fmt.Printf("%d \t %b \t %x \t %#X \n", 42, 42, 42, 42)

	//We can also convert a number to its UTF-8 equivalent
	fmt.Printf("%q \n", 42)

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

	//var asd1, dsa2 = "this is", 2

	//short variables are a short way to initialize variables
	short := "short"
	shortb := true

	//you can initialize many variables of dif types at once
	boolvar, intvar, floatvar := true, 64, 64.2

	//with single quotes it gets converted to unicode number
	unicode := 'm'

	backticks :=
		`
	with backticks
	you dont need to escape endlines or tabs
	or "quotes"`

	fmt.Println(a, text, number, text2, asd, dsa, boolean, short, shortb, boolvar, intvar, floatvar, unicode, backticks)

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