package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algobra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func toposort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// When an anoymous function requires recursion, we must first declare a variable,
	// and then assign the anonymouse function to that variable. Had these two steps been
	// combined in the declaration, like
	//
	// 		visitAll := func(item []string) {
	// 			// ...
	// 			visitAll(m[item]) // compile error: undefined: visitAll
	// 			// ...
	// 		}
	//
	// the funciton literal would not be within the scope of the variable visitAll
	// so it would have no way to all itself recursively
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			// depth first traversal
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m { // value can be omitted
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order

}