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

	//short variables dont need var keyword and are local variables (kind of let in JS)
	short := "short"
	shortb := true

	fmt.Println(a, text, number, text2, asd, dsa, boolean, short, shortb)

	const CONSTANT = "GOPHER"
}

const EXAMPLE = "You can declare const and vars outside of functions, if they are not short variables"
