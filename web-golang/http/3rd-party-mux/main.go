package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//We can use 3rd party mux like this one which has different naming
	mux := httprouter.New()
	mux.GET("/", index)
	//here we dont care about the last slash, we dont need to use it to have subroutes
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	//we can handle variables in the URL easier (go to func user or blog***)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)

	//in the end we always use listenandserve but we use different mux if we want
	http.ListenAndServe(":8080", mux)
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//we get the URL params as 3rd parameter and then we can search ByName
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		//replies with error message and code
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
