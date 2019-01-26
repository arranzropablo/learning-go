package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	//In Go only vars and funcs starting with capital letter are exported, those with lower letter are "private"
	//The body is type ReadCloser interface which implements Reader and Closer interfaces (we have to implement Read and Close functions to imlpement those interfaces)
	response, err := http.Get("http://www.google.com")
	if err == nil {

		//
/*		//We initialize it with this number bc the Read function dont resize the [], just copy the body into []byte and if it is full it will fail
		bs := make([]byte, 99999)
		//We pass bs to Read which fills it (return num of bytes and error, not a byte slice)
		response.Body.Read(bs)
		fmt.Println(string(bs))*/

		//instead of all that Read stuff we can directly copy the body to stdout (io.Copy uses a Writer and a Reader)
		//io.Copy(os.Stdout, response.Body)

		//We can use our own type if it implements the interface Writer
		io.Copy(logWriter{}, response.Body)

	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}