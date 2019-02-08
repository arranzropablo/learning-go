package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8888", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")

	if err != nil {
		c = &http.Cookie{
			Name: "counter",
			Value: strconv.Itoa(0),
			Path: "/",
		}
	} else {
		counter, _ := strconv.Atoi(c.Value)
		c.Value = strconv.Itoa(counter + 1)
	}

	http.SetCookie(w, c)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counter")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}