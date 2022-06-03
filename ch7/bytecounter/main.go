package main

import (
	"fmt"
)

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	// The conversion is required to make the types of len(p) and *c match in the
	// += assignment statement
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}