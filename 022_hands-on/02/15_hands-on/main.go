package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	//"os"
	//"strings"
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
			continue
		}

		//io.WriteString(conn, "Hello world!\n")
		//io.WriteString(conn, fmt.Sprintf("Content-Length: %v\n", len("Hello world!\n")))
		//conn.Close()
		//io.WriteString(os.Stdout, "Hello world!\n")
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	getRequests(conn)
	fmt.Println("getRequests: CALLED")
	texts := fmt.Sprintln("<!DOCTYPE html><html lang=\"en\"><head><title>My Name is Anthony Gozalvez</title></head><body><h1>HOLY COW THIS IS LOW LEVEL</h1></body></html>")
	writeResponse(conn, texts)
	defer conn.Close()

	//io.WriteString(conn, "I am there.")
	fmt.Println("Code got here.")
	//io.WriteString(conn, "I see you connected.")
}

func getRequests(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		//reqT := strings.Fields(line)[0]
		//reqURI := strings.Fields(line)[1]
		fmt.Println(line)

		if line == "" { break; }
		//line1 := strings.Fields(line)[0]
		//return line1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner Scan Error", err)
	}

}


func writeResponse(conn net.Conn, texts string) {
	//fmt.Println(texts)
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(texts))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, texts)
}

