package main

import "fmt"

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


}