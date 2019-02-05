package main

import (
	"net/http"
)

//we dont need listenandserve or main (we just use init) when deploying go1.9
func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
}
