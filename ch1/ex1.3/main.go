package main

import (
	"fmt"
	"time"
	"os"
	"strings"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
	    s += sep + arg
	    sep = " "
	}
	fmt.Println(s)
	fmt.Printf("time elasped using +: %f\n", time.Since(start).Seconds()) // 0.000085

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("time elasped using strings.Join: %f\n", time.Since(start).Seconds()) // 0.000006
}