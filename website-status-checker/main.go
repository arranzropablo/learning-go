package main

import (
	"fmt"
	"net/http"
	"time"
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

	//c is a channel for strings
	c:= make(chan string)

	for _, link := range links {
		//with go instruction, it executes a call in a different goroutine
		//the go scheduler runs one routine until it finishes or makes a blocking call
		//By default Go tries to use one core but each CPU can run one go routine
		//By default we have the main routine and if we use the keyword "go" we create child routines
		go checkLink(link, c)
		//If we don't wait for the children routines to finish the main routine will finish its work and exit without fetching the other routines
		//We have channels to comunicate between goroutines, we send data to a channel and it gets sent to the other routines
		//Channels are typed (the data inside them must be the same type)
		//We create channles to check if routines have finished and let outher routines continue
	}

	//for i := 0; i < len(links); i++{
	//	fmt.Println(<-c)
	//}

	//this makes an infinite loop
	//for {
	//
	//}

	for item := range c {
		//if we put this sleep here we are pausing 5 seconds our main routine
		//therefore we are waiting 5 seconds between every request
		//time.Sleep(5 * time.Second)

		//To wait 5 seconds in each separated routine we place the sleep in a function literal (anon func)
		//If we don't pass the item, the var item is declared in the outter scope of the func and will do weird stuff (requests to the same url repeatedly)
		//We should never use the same var in different routines, we should pass the vars (they are passed by value)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(item)
	}
}

//Concurrency vs Parallelism
//Concurrency: If one thread blocks another thread is picked up (but they don't execute simultaneously)
//Parallelism: We can execute multiple threads at the same time (multiple CPUs required)

//func checkLink(link string, c chan string) {
//	_, err := http.Get(link)
//	if err != nil{
//		c<-"Error during request to " + link
//	} else {
//		c<-link + " is up!"
//	}
//}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link, "is down!")
		c<-link
	} else {
		fmt.Println(link, "is up!")
		c<-link
	}

	//If we put this sleep here it's useless because it's after the c<-link and it will sleep after sending and we won't notice the pause
	//time.Sleep(5 * time.Second)
}