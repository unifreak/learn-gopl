// Echo4 prints its command-lin arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// Pointer are key to flag package.
// `flag.Bool(name, default value, help message)` create flag variables.
// `n` is a pointer to it, so must be accessed indirectly as *n. See below.
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // update flag variable from their default values.
	fmt.Printf(strings.Join(flag.Args(), *sep)) // `flag.Args()` contain non-flag args
	if !*n {
		fmt.Println()
	}
}