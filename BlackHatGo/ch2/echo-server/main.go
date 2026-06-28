// echo-server from Black Hat Go

package main

import (
	"io" // missing from the book the program would compile beacuse of io.EOF being called
	"log"
	"net"
)

// echo is a handler function that simply echoes recieved data.
func echo(conn net.Conn) {
	defer conn.Close()
	
	// Create a buffer to store recieved data.
	b := make([]byte, 512)
	for {
		// Recieved data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Send data via conn.Write
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// Bind to tcp port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for conneciton. Create net.Conn on connection established
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unalbe to accept connection")
		}
		// Handle the connection, using goroutine for concurrency.
		go echo(conn)
	}
}
