package main

import "fmt"

type bot interface{
	getGreeting() string

	//We don't add printGreeting to our interface because it will have the same code for both bots
	//Instead, we pass a type bot as parameter
	//printGreeting(b bot)
}

type englishBot struct{}

type spanishBot struct{}

//if we are not making use of the value we can omit the var and just write the type
func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}