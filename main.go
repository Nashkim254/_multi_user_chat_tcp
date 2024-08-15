package main

import (
	"log"
	"net"
)

const port = "6969"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message := []byte("Hello user")
	n, err := conn.Write(message)
	if err != nil {
		log.Printf("Could not write message to %s : %s\n", conn.RemoteAddr(), err)
		return
	}
	if n < len(message) {
		log.Printf("The message was not fully written %d %d\n", n, len(message))
	}
}
func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error connecting to: %s on %s", port, err)
	}
	log.Printf("Listening to TCP connections on port %s ....\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not accept connection %s\n", err)
		}
		go handleConnection(conn)
	}
}
