package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// lister to port
	srv, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		connection, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		// handle the connect
		go handleServerConnection(connection)
	}
}

func handleServerConnection(conn net.Conn) {
	// Receive the message
	var msg string
	fmt.Println("literal connection", conn)
	err := gob.NewDecoder(conn).Decode(&msg)

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Received", msg)
	}
}

func client() {
	// connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "fuck the world"
	fmt.Println("sending", msg, conn)
	err = gob.NewEncoder(conn).Encode(msg)

	if err != nil {
		fmt.Println(err)
	}

	conn.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}