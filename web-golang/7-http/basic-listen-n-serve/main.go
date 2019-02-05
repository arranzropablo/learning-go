package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//any type we want can be a handler if it implements ServeHTTP
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//there's a TCP server under this and if we see the implementation is the same we know
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	//in this case we just listen on 8888 and serve just that println
	http.ListenAndServe(":8888", d)
}
