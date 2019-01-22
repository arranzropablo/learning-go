package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	//In Go only vars and funcs starting with capital letter are exported, those with lower letter are "private"
	response, err := http.Get("http://www.google.com")
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		newStr := buf.String()
		fmt.Println(newStr)
	}



}