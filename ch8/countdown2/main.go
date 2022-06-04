package main

import (
	"time"
	"fmt"
	"os"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// We can't just receive from each channel (abort and tick) because whichever
	// operation we try first will block until completion. Wee need to multiplex these
	// operations, and to do that, we need a select statement.
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	// The time.After function immediately return a channel, and starts a new
	// goroutine that sends a single value on that channel after the specified time.
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Launched.")
}