// Add types, constants, and functions to tempconv for processing temperatures
// in the Kelvin scale, where zero Kelvin is -273.15C and a difference of 1K has
// the same magnitude as 1C.
package main

import (
	"os"
	"fmt"
	"strconv"
	"learn/gopl/ch2/ex2.1/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.1: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.KToC(k), c, tempconv.CToK(c))
	}
}