// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"net"
	"time"
	"log"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // block until incoming connection request is made
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time. Sequential.
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}