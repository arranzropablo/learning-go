package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}


func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Hello you are connected\n")
	scan := bufio.NewScanner(conn)

	for scan.Scan() {
		io.WriteString(os.Stdout, scan.Text())
	}
	io.WriteString(conn, "Code never got here\n")

}
