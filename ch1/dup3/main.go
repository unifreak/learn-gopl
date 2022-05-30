// Dup3 read from file line by line and print duplicated lines count
package main

import (
	"fmt"
	// As of Go 1.16, the same functionality is now provided by package io or
	// package os, and those implementations should be preferred in new code
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// `.ReadFile` return a byte slice that must be converted into string
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}