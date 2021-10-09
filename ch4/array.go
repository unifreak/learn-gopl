package main

import (
	"fmt"
)

func main() {
	var a [3]int 				// array of 3 integers
	fmt.Println(a[0]) 			// print the first element
	fmt.Println(a[len(a)-1]) 	// print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// Initialize
	// 1. with array literal
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	// use ... to determin length by the number of initializers.
	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	// size of array is part of its type
    // size must be a constant expression (can be computed when compiled)
    // q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

    // 2. specify index and value pairs
    type Currency int
    const (
    	USD Currency = iota
    	EUR
    	GBP
    	RMB
    )
    symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
    fmt.Println(RMB, symbol[RMB]) // "3 ¥"

    // indices can appear in any order or omitted
    p := [...]int{99: -1}

    // Comparable
    // if element type is comparable then array type is comparable, too.
    a := [2]int{1, 2}
    b := [...]intt{1, 2}
    c := [2]int{1, 3}
    fmt.Println(a == b, a == c, b == c) // "true false false"
    d := [3]int{1, 2}
    fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}