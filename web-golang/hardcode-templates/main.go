package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main (){
	//We can hardcode an HTML template in a string and replace the vars we want, like name
	//We can take this name from many places like command line arg or Scan or wherever
	name := "Pablo"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	//We defer closing the file until the func has finished
	defer file.Close()

	io.Copy(file, strings.NewReader(tpl))

}
