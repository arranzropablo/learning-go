package main

//you can call functions from other files if they are in the same package and no need to import
import "fmt"

func main() {
	fmt.Println(getAOfClubs())
	deckSize := deckSize()
	fmt.Println(deckSize)
}

//functions can return stuff if return types are specified between parentheses (if more than one)
func getAOfClubs() (string, string) {
	return "Ace", "Clubs"
}

func deckSize() int64 {
	return 52
}