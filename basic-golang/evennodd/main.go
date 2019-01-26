package main

import "fmt"

func main() {
	slice := []int{0,1,2,3,4,5,6,7,8,9,10}

	//you can also make a slice with:
	//greeting := make([]string, 3, 5)
	//the initial length will be 3 and the capacity will be 5 meaning you wont be able to assign a value for a pos greater than 2
	//because it does not exist at the moment
	//using append you can go over the capacity and/or length
	for _, n := range slice {
		if n % 2 == 0{
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}