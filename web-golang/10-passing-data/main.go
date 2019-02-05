package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", val)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func val(w http.ResponseWriter, req *http.Request) {
	//this can retrieve params from forms or URLs (but body params have preference)
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

// visit this page:
// http://localhost:8080/?q=dog
