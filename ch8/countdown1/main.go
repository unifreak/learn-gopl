package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("Commencing countdown.")

	// The time.Tick function returns a channel on which it sends events
	// periodically, acting like a metronome.
	//
	// The Tick function is convenient, but it's appropriate only when the ticks will
	// be needed throughtout the lifetime of the application.
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Luanched.")
}