package main

import (
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip int64
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
//array de hoteles por region

	h := []region{
		{
			Region: "Southern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
		{
			Region: "Northern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
		{
			Region: "Central",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     95612,
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     95612,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h)
	if err != nil {
		panic(err)
	}
}
