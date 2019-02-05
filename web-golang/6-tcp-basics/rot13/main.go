package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	//This way we listen to telnet messages in that connection
	li, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	//we defer the closing til main ends
	defer li.Close()

	for {
		//We accept every connection request (telnet localhost 8888)
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		//We start a goroutine to handle that connection and avoid blocking the app while it's handled
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//We can set a deadline for the connection, and it will be closed automatically after certain time
	err := conn.SetDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	//Scanning input from connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)
		fmt.Fprintf(conn, "%s - %s \n\n", ln, r)
	}

	//This part of code is never reached unless the connection is closed by the client (because scanner will be scanning otherwise)
	defer conn.Close()
	fmt.Println("Close connection by client")
}

func rot13(bs []byte) []byte{
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
