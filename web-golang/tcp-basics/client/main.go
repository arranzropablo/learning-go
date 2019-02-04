package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//We can create a client to call tcp connection and read or write
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//read stuff from tcp server
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
	//send info to tcp server
	fmt.Fprintln(conn, "I dialed you.")
}
