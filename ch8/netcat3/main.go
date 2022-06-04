package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Comparing to netcat2, we use a channel to synchronize two goroutine.
//
// When the user closes the standard input stream, mustCopy returns and the main
// goroutine calls conn.Close(), closing both halves of the network connection.
// Closing the write half of the connection causes the server to see an end-of-file
// condition. Closing the read half causes the background goroutine's call to
// io.Copy to return a "read from closed connection" error, which is why we've
// removed the error logging.
//
// Before it returns, the background goroutine logs a message, then sends a value on
// the done channel. The main goroutine waits until it has received this value before
// returning. As a result, the program always logs the "done" message before exiting.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	// Notice the go statement calls a literal function, a common construction
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutin
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done 	// wait for background gorountine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}