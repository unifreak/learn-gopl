package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	// Go programmers often describe a map used in this fashion as `set of strings`
	// But beware, not all map[string]bool are simple sets; some may contain
	// both true and false values.
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}