// Ex2.2 is a general-purpose unit-conversion program
package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
	"./ex2.2"
)

var t = flag.Bool("t", false, "convert temperatures.")
var l = flag.Bool("l", false, "convert lengths.")
var w = flag.Bool("w", false, "convert weights.")

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
			os.Exit(1)
		}

		if *t {
			f := conv.Fahrenheit(n)
			c := conv.Celsius(n)
			fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
		}
		if *l {
			ft := conv.Feet(n)
			m := conv.Meter(n)
			fmt.Printf("%s = %s, %s = %s\n", ft, conv.FtToM(ft), m, conv.MToFt(m))
		}
		if *w {
			lb := conv.Pound(n)
			kg := conv.Kilogram(n)
			fmt.Printf("%s = %s, %s = %s\n", lb, conv.LbToKg(lb), kg, conv.KgToLb(kg))
		}
	}
}