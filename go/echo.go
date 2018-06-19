package main

import (
    "bufio"
    "fmt"
    "strings"
	"log"
	"net"
	"time"
)

// go run echo.go &
// go run netcatEcho.go
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

// E.g. hello, echoes to HELLO, Hello, hello
func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// Prints echo to client
func handleConn(conn net.Conn) {
    // Read from connection
    input := bufio.NewScanner(conn)
    for input.Scan() {
        echo(conn, input.Text(), 1*time.Second)
    }
    conn.Close()
}
