package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	//NewScanner returns a new Scanner to read from r.
	// The split function defaults to scanlines
	scanner := bufio.NewScanner(conn)
	//Scann advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method.
	for scanner.Scan() {
		//Text returns the most recent token generated by a call to Scan
		// as a newly allocated string holding its bytes/
		ln := scanner.Text()
		fmt.Println(ln)
		ln = fmt.Sprint("FROM SERVER: " + ln)
		fmt.Fprintln(conn, ln)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		//only handles a single connection at once
		//io.Copy(conn, conn)
		handle(conn)
		conn.Close()
	}
}

//run this then,
// telnet localhost 9000
