package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	//The dot gets replaced by the data we are passing, if we want to have "fields" we can pass a struct or map
	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Pablo")
	//if err == nil {
	//	log.Fatalln(err)
	//}

	//err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", []string{"Pablo", "Adri", "Javi"})
	//if err == nil {
	//	log.Fatalln(err)
	//}

	//sages := map[string]string{
	//	"India":    "Gandhi",
	//	"America":  "MLK",
	//	"Meditate": "Buddha",
	//	"Love":     "Jesus",
	//	"Prophet":  "Muhammad",
	//}
	//err := tpl.ExecuteTemplate(os.Stdout, "map.gohtml", sages)
	//if err == nil {
	//	log.Fatalln(err)
	//}

	type sage struct {
		Name  string
		Motto string
	}
	err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", sage{"Buddha", "The belief of no beliefs"})
	if err == nil {
		log.Fatalln(err)
	}
}