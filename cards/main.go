package main

//you can call functions from other files if they are in the same package and no need to import
import (
	"fmt"
)

func main() {
	fmt.Println(getAOfClubs())
	deckSize := deckSize()
	fmt.Println(deckSize)

	suites := []string{"Hearts", "Clubs", "Spades"}

	//append, delete, etc... does not modify the object, but return a copy (we can reassign then)
	suites = append(suites, "Diamonds")
	fmt.Println(suites)

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