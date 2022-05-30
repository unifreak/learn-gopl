// Outline print the structure of the html doc's node tree in outline.
//
// Usage: ./fetch https://golang.org | ./outline
package main

import (
	"os"
	"fmt"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

// Note that although `outline` push an element on stack, there is no
// corresponding pop. When outline calls itself recursively, the callee
// receives a copy of stack. Although the callee may append elements
// to this slice, modifying its underlying array and perhaps even allocating
// a new array, it doesn't modify the initial elements that are visible
// to the caller, so when the function returns, the caller's stack is as
// it was before the call.
//
// Many programming language implementations use a fixed-size function call
// stack; size from 64KB to 2MB are typical. Fixed-size stacks impose a
// limit on the depth of recursion, son one must be careful to avoid a
// stack overflow when tranversing large data structures recursively;
// fixed-size stacks may even pose a security risk.
//
// In contrast, typical Go implementations use variable-size stacks that
// start small and grow as needed up to a limit on the order of a gigabyte.
// This lets us use recursion safely and without worrying about overflow.
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}