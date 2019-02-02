package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	//we handle this route with index func
	http.HandleFunc("/", index)

	//we handle public route with the public directory
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	//we listen in 8080 port
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	//we execute the template in the responsewriter
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
