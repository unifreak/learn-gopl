// Dup1 prints the text of each line that appears more than
// once in the stardard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// `map` holds kv pairs, key can be any type can be compared
	// with ==. constant-time operation
	counts := make(map[string]int) // make() create new map
	input := bufio.NewScanner(os.Stdin)
	// each call to .Scan() read next line and remove tailing newline char.
	// result can be retrived by .Text()
	// Scan function returns true if there is a line and false if there is no more input.
	for input.Scan() {
		counts[input.Text()]++
	}
	// common `verbs`:
	// %d 				decimal
	// %x, %o, %b 		hex, octal, binary
	// %f, %g, %e 		3.141593, 3.141592653589793, 3.141593e+00
	// %t 				true or false
	// %c 				rune
	// %s 				string
	// %q 				quoted string "abc" or rune 'c'
	// %v 				natual format
	// %T 				type of any value
	// %% 				literal %

	// NOTE: ignoring potential errors from input.Err()

	// The order of map iteration is not specified, but in practice it is random
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}