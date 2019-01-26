package main

import (
	"fmt"
)

func main() {
	//This is how you define a map
	m := map[string]string{
		"initialkey1": "initialvalue1",
		"initialkey2": "initialvalue2",
		}

	m["keytype"] = "valuetype"

	fmt.Println(m)

	//Other way to create an empty map
	var m2 map[string]string
	fmt.Println(m2)

	//Other way to create an empty map
	m3 := make(map[string]string)
	fmt.Println(m3)

	delete(m, "keytype")
	fmt.Println(m)

	colors := map[string]string{
		"black": "#000000",
		"white": "#ffffff",
	}

	for color, hex := range colors {
		fmt.Println(color, hex)
	}

	name := "Pablo"

	switch name {
		//We can evaluate more than one case
		case "Pablo", "Someone":
			fmt.Println("Hello")
			//fallthrough keyword executes next case also
			fallthrough
		case "":
			fmt.Println("Who are you")
		default:
			fmt.Println("Default")
	}

	//We can also make a switch with no condition and it acts like if-else statement
	switch {
		case name == "Pablo":
			fmt.Println("Hello Pablo")
		case len(name) >= 1:
			fmt.Println("Expected length")
		default:
			fmt.Println("wait what")
	}

	//We can also switch over the type of a var
	SwitchOnType(name)
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}