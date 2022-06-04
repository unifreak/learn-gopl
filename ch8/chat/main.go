// Chat is a server that lets clients chat with each other.
//
// There are four kinds of goroutine in this program. There is one instance apiece of
// the main and broadcaster goroutines, and for each client connection there is one
// handleConn and one clientWriter goroutine. The broadcaster is a good illustration
// of how select is used, since it has to respond to three different kinds of
// messages.
//
// While hosting a chat session for n clients, this program runs 2n+2 concurrently
// communicating goroutines, yet it needs no explicit locking operations. The
// clients map is con- fined to a single goroutine, the broadcaster, so it cannot be
// accessed concurrently. The only variables that are shared by multiple goroutines
// are channels and instances of net.Conn, both of which are concurrency safe.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// The broadcaster listens on the global entering and leaving channels for
// announcements of arriving and departing clients. When it receives one of these
// events, it updates the clients set, and if the event was a departure, it closes
// the client's outgoing message channel. The broadcaster also listens for events on
// the global messages channel, to which each client sends all its incoming
// messages. When the broadcaster receives one of these events, it broadcasts the
// message to every connected client.

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// The handleConn function creates a new outgoing message channel for its client and
// announces the arrival of this client to the broadcaster over the entering
// channel. Then it reads every line of text from the client, sending each line to
// the broadcaster over the global incoming message channel, prefixing each message
// with the identity of its sender. Once there is nothing more to read from the
// client, handleConn announces the departure of the client over the leaving channel
// and closes the connection.
//
// In addition, handleConn creates a clientWriter goroutine for each client that
// receives messages broadcast to the client’s outgoing message channel and writes
// them to the client's network connection. The client writer’s loop terminates
// when the broadcaster closes the channel after receiving a leaving notification.
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

// The job of the main goroutine, shown below, is to listen for and accept incoming
// network connections from clients.
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}