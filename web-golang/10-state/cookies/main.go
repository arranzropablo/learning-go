package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	//we set a cookie with the responsewriter and a pointer to a cookie
	//we can hava multiple cookies
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path: "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	//we can retrieve a cookie like this
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

func delete(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//to delete a cookie in a certain time we can set maxage, we can also set maxage lower than 0 to delete it right now, or 0 to never be deleted
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Fprintln(w, `<h1><a href="/set">Cookie deleted. set a cookie</a></h1>`)
}
