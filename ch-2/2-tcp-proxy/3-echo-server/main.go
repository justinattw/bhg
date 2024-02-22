package main

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)

	// Indefinite loop
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err != nil && err != io.EOF {
			log.Println("Unexpected error")
			break
		}

		if err == io.EOF && size == 0 {
			log.Println("Client disconnected")
			break
		}

		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	// Infinite loop ensures server will continue to listen for connections even after one has been received
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept() // This function blocks execution as it awaits client connections
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
