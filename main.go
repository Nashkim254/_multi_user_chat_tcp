package main

import (
	"log"
	"net"
)

const port = "6969"
const StreamerMode = false

type Client struct {
	conn   net.Conn
	output chan string
}

func SafeRemoteAddr(conn net.Conn) string {

	if StreamerMode {
		return "[REDUCTED]"
	} else {
		return conn.RemoteAddr().String()
	}

}

func handleConnection(conn net.Conn, outgoing chan string) {
	defer conn.Close()

	buffer := []byte{}
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			return
		}
		outgoing <- string(buffer[0:n])
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
		log.Printf("Accepting connection from %s", SafeRemoteAddr(conn))
		outgoing := make(chan string)
		go handleConnection(conn, outgoing)
	}
}
