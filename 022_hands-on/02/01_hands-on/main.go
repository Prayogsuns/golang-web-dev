package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)


func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Fatal err on Listen: ", err)
	}
	defer li.Close()

        for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Connection error on Accept: ", err)
		}

		io.WriteString(conn, "Hello world!\n")
		io.WriteString(conn, fmt.Sprintf("Content-Length: %v\n", len("Hello world!\n")))
		conn.Close()
		io.WriteString(os.Stdout, "Hello world!\n")
	}
}
