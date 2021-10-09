package main

// key type of graph is string, and the value type is map[string]bool
// representing a set of strings. Conceptually, graph maps a string
// to a set of related strings, its successors in a directed graph
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}