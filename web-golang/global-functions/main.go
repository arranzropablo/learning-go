package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	//functions already existing
	//or, and, eq

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err = tpl.ExecuteTemplate(os.Stdout, "and.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "comparison.gohtml", g1)
	if err != nil {
		log.Fatalln(err)
	}

	//another global function is template, with it you can define templates and use them in other templates
	//in the "subtemplate":
	//{{define <name>}}code{{end}}
	//in the one "outside":
	//{{template <name>}}
}
