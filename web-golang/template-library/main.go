package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	//With template library we can parse html files
	//the extension for templates is gohtml as a convention, but we can have any extension we want
	template, err := template.ParseFiles("tpl.gohtml")
	//file, err := os.Open("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	//This applies data (nil) to the template and puts it in the desired output (stdout in this case)
	template.Execute(os.Stdout, nil)

	//We can also create a new file and use it as the desired output (it implements the writer interface)
	//nf, err := os.Create("newfile.gohtml")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//template.Execute(nf, nil)

	//an object template is kind of a container of templates, this way we can add more templates to that container:
	template, err = template.ParseFiles("one.gmao", "two.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	//and then we can execute a single template with
	err = template.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//As we dont want to type every filename, most typical func to use is:
	template, err = template.ParseGlob("templates/*")
	err = template.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

//We only want to parse our templates once, so a good practice is parsing them on init func
//init func can be declared more than once in the same package or file and it get executed in order of appearance
//init can't be referenced in other functions

//var tpl *template.Template
//func init() {
//	"Must" panics if the error is not null (used to do error checking when working with templates)
//	tpl = template.Must(template.ParseGlob("templates/*"))
//}
