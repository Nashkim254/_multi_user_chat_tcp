package main

import (
	"log"
	"net"
)

const port = "6969"
const StreamerMode = false

func SafeRemoteAddr(conn net.Conn) string {

	if StreamerMode {
		return "[REDUCTED]"
	} else {
		return conn.RemoteAddr().String()
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message := []byte("Hello user\n")
	n, err := conn.Write(message)
	if err != nil {
		log.Printf("Could not write message to %s : %s\n", SafeRemoteAddr(conn), err)
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
		log.Printf("Accepting connection from %s", SafeRemoteAddr(conn))
		go handleConnection(conn)
	}
}
