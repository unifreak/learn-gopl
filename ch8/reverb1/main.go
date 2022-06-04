package main

import (
	"fmt"
	"strings"
	"bufio"
	"time"
	"net"
	"log"
)

// Usage
// 		$ go build ch8/reverb1
// 		$ ./reverb1 &
// 		$ go build ch8/netcat2
// 		$ ./netcat2
// 		Hello?
// 		 	HELLO?
// 		 	Hello?
// 		 	hello?
// 		Is there anybody there?
// 			IS THERE ANYBODY THERE?
// 		Yoo-hooo!
// 			Is there anybody there?
// 			is there anybody there?
// 			YOO-HOOO!
// 			Yoo-hooo!
// 			yoo-hooo!
//
// Notice that the third shout from the client is not dealt with until the second shout
// has petered out, which is not very realistic. See reverb2.
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// The arguments to the function started by go are evaluated when the go statement
		// itself is executed; thus input.Text() is evaluated in the main goroutine.
		echo (c, input.Text(), 1 * time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}