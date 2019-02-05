package main

import (
	"io"
	"net/http"
)

//type hotdog int
//
//func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	io.WriteString(res, "dog dog dog")
//}
//
//type hotcat int
//
//func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	io.WriteString(res, "cat cat cat")
//}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	//var d hotdog
	//var c hotcat
	//
	//We can create a new serve mux and pass it as a handler to ListenAndServe or we can use default serve mux passing nil
	//mux := http.NewServeMux()
	//
	//The slash after dog is to allow receiving /dog/something/else... in case of /cat if we call /cat/something we will get an error
	//http.Handle("/dog/", d)
	//http.Handle("/cat", c)

	//A better way to handle paths (func instead of a Handler)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8888", nil)
}
