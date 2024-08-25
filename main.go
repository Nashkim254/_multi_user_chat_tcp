package main

import (
	"fmt"
	"log"
	"net"
)

const port = "6969"
const StreamerMode = false

type MessageType int

const (
	ClientConnected MessageType = iota + 1
	DeleteClient
	NewMessage
)

type Message struct {
	Type MessageType
	Conn net.Conn
	Text string
}
type Client struct {
	conn   net.Conn
	output chan string
}

func server(messages chan Message) {
	conns := map[string]net.Conn{}
	for {
		msg := <-messages
		switch msg.Type {
		case ClientConnected:
			conns[msg.Conn.RemoteAddr().String()] = msg.Conn
		case DeleteClient:
			delete(conns, msg.Conn.RemoteAddr().String())
			msg.Conn.Close()
		case NewMessage:
			for _, conn := range conns {
				_, err := conn.Write([]byte(msg.Text))
				if err != nil {
					fmt.Printf("Could not send data, %s: %s", SafeRemoteAddr(conn), err)
				}
			}
		}
	}
}

func SafeRemoteAddr(conn net.Conn) string {

	if StreamerMode {
		return "[REDUCTED]"
	} else {
		return conn.RemoteAddr().String()
	}

}

func client(conn net.Conn, messages chan Message) {
	defer conn.Close()

	buffer := make([]byte, 512)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			messages <- Message{
				Type: DeleteClient,
				Text: string(buffer[0:n]),
				Conn: conn,
			}
			return
		}
		messages <- Message{
			Type: NewMessage,
			Text: string(buffer[0:n]),
			Conn: conn,
		}
	}
}
func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error connecting to: %s on %s", port, err)
	}
	log.Printf("Listening to TCP connections on port %s ....\n", port)
	messages := make(chan Message)
	go server(messages)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not accept connection %s\n", err)
		}
		log.Printf("Accepting connection from %s", SafeRemoteAddr(conn))
		messages <- Message{
			Type: ClientConnected,
			Conn: conn,
		}
		go client(conn, messages)
	}
}
