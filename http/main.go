package main

import (
	"fmt"
	"net/http"
)

func main() {
	//In Go only vars and funcs starting with capital letter are exported, those with lower letter are "private"
	//The body is type ReadCloser interface which implements Reader and Closer interfaces (we have to implement Read and Close functions to imlpement those interfaces)
	response, err := http.Get("http://www.google.com")
	if err == nil {

		//We initialize it with this number bc the Read function dont resize the [], just copy the body into []byte and if it is full it will fail
		bs := make([]byte, 99999)
		//We pass bs to Read which fills it (return num of bytes and error, not a byte slice)
		response.Body.Read(bs)
		fmt.Println(string(bs))

		//We don't read like that, Go has helpers, one way found on internet is:
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(response.Body)
		//newStr := buf.String()
		//fmt.Println(newStr)
	}



}