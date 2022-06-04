package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)

	// Here we copy input to the server in main goroutine, so the programe terminates
	// as soon as the input stream closes, even if the background goroutine is still
	// working. To make the program wait for the background gorouine to complete before
	// exiting, we can use a channel to synchronize the two gorouines. See netcat3.
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}