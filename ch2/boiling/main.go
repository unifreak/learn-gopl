// Boiling prints the boiling point of water.
package main

import "fmt"

// declared outside of a function, boilingF is visible in all files of package
// to which it belongs. if name begins with an upper-case letter, it is `exported`,
// visible and accessible outside of tis own package
const boilingF = 212.0  // @? why not := ?
						// @A short declaratioin is only available inside func.

func main() {
	// declared inside a function, f and c is local to function main()
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}