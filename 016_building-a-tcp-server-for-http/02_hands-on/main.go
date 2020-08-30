package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln[0:3] == "GET" {
			urlP := strings.Split(ln, " ")
			urlStr := strings.Join(urlP[1:len(urlP)-1], " ")
			fmt.Fprintf(conn, "%s\n", conn.LocalAddr().String()+urlStr)
		}
		break
	}
	defer conn.Close()

	fmt.Println("Test print")
}
