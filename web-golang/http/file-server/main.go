package main

import (
	"io"
	"net/http"
)

func main() {
	//single file with io.Copy, ServeContent and ServeFile
	//this handles root URL showing the files in this dir
	http.Handle("/", http.FileServer(http.Dir(".")))

	//This way when we call from img tag resources/toby.jpg we are removing the resources part and searching the
	//result in ./assets
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/dog", dog)

	//NotFoundHandler returns a 404 page not found, we can use it for example for favicon
	//http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

//if we want a static file server (static page or whatever) we can just have this in the main
//func main() {
//	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
//}
//and then have a file named index.html which will automatically render


func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//when calling toby.jpg it goes to localhost:8080/toby.jpg and there it is the image
	//io.WriteString(w, `<img src="toby.jpg">`)
	io.WriteString(w, `<img src="resources/toby.jpg">`)
}