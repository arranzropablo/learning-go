package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	//mux := http.NewServeMux()
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page!")
	}))
	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Woof WOof!")
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pablo Arranz Ropero")
	})
	http.HandleFunc("/template/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "template.gohtml", "Pablo")
	})
	http.ListenAndServe(":8080", nil)
}
