package main

import (
	"log"
	"net"
)

const port = "6969"

func handleConnection(conn net.Conn) {

}
func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error connecting to: %s on %s", port, err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not accept connection %s", err)
		}
		go handleConnection(conn)
	}
}
