package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			method := strings.Fields(ln)[0] //divided by whitespace
			fmt.Println("WE PRINTED THIS - METHOD : ", method)
		} else {
			if ln == "" {
				break
			}
		}
		i++
	}

	//response
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<strong>Hello World</strong>
	</body>
	</html>	
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")            //status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) //headers with another option for writng to connection
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func main() {
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handleConn(conn) // what does go do?? threading?
	}
}
