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

// Conceptually, the prerequisite information forms a directed graph with a node
// for each course and edge from each course to the courses that it depends on.
// The graph is acyclic: there is no path from a course that leads back to itself.
// We can compute a valid sequence using *depth-first search* through the graph.
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
	// so it would have no way to call itself recursively
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