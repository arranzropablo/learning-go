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

	//this takes too long because http.Get is a blocking call
	//for _, link := range links {
	//	if _, err := http.Get(link); err != nil{
	//		fmt.Println("Error during request to", link)
	//	} else {
	//		fmt.Println(link, "is up!")
	//	}
	//}

	for _, link := range links {
		//with go instruction, it executes a call in a different goroutine
		//the go scheduler runs one routine until it finishes or makes a blocking call
		//By default Go tries to use one core but each CPU can run one go routine
		//By default we have the main routine and if we use the keyword "go" we create child routines
		go checkLink(link)
	}
}

//Concurrency vs Parallelism
//Concurrency: If one thread blocks another thread is picked up (but they don't execute simultaneously)
//Parallelism: We can execute multiple threads at the same time (multiple CPUs required)

func checkLink(link string) {
	if _, err := http.Get(link); err != nil{
		fmt.Println("Error during request to", link)
	} else {
		fmt.Println(link, "is up!")
	}
}