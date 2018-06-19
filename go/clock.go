package main

import (
	"io"
	"log"
	"net"
	"time"
)

// go run clock.go &
// nc localhost 8000
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		// Abort connection, continue
		if err != nil {
			log.Print(err)
			continue
		}
        // Spawn thread to handle connection
		go handleConn(conn)
	}
}

// Writes the time infinitely per second to accepted connection
func handleConn(conn net.Conn) {
	// Defer the closing of connection
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		// Client disconnect, etc.
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
