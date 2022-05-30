package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack() // for diagnostic purpose
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}