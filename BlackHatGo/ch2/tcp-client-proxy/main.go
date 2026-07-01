package main

import (
	"net"
	"log"
	"io"
)

// create proxy from BHGO example

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "192.168.1.104:8080") // local box with python doing simple webserver
	if err != nil {
		log.Fatalln("Unable to connect to website")
	}
	defer dst.Close()
	
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}() // call the go func
	
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
		}
}

func main() {
	listener, err := net.Listen("tcp", ":9100") // looks like most ports used by printers 
												// might get lost in the clutter
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}

