package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		if _, err := http.Get(link); err != nil{
			fmt.Println("Error during request to", link)
		} else {
			fmt.Println(link, "is up!")
		}
	}
}