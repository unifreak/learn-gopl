package main

import (
	"fmt"
    "math"
	// "time"
	// "os"
	// "io"
)

func incr(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r)
}


func change(s []byte) []byte {
    s[0] = 'b';
    return s
}

func main() {
    for _, r := range "Hello, 世界" {
        fmt.Printf("%t", r)
    }

    s := []string{"a", "b"}
    fmt.Printf("%q\n", s)
    fmt.Printf("%v\n", []byte("a"))
}