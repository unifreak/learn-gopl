package main

import (
	"fmt"
)

// Finite pipeline indicated by closing a channel.
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		// equivalent to:
		//
		// 		for {
		// 			x, ok := <-naturals
		// 			if !ok {
		// 				break
		// 			}
		// 			squares <- x*x
		// 		}
		//
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}