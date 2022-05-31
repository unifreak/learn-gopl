package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q")) // ""
	fmt.Println(m.Get("item")) // "1"  (first value)
	fmt.Println(m["item"]) // "[1 2]" (direct map access)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3") // panic: assignment to entry in nil map

	// Because url.VAlues is a mpa type and a map refers to its key/value pairs
	// indireclty, any updates and deletions that url.Values.Add makes to the
	// map elements are visible to the caller. However, as with ordinary funcitons,
	// any changes a method makes to the reference itself, like setting it to nil
	// or making it refer to a differenct map data structure, will not be reflected
	// in the caller.
}