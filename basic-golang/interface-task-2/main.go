package main

import (
	"io"
	"os"
)

//go run main.go myfile.txt
func main(){
	f, err := os.Open(os.Args[1])

	if err == nil {
		io.Copy(os.Stdout, f)
	}
}