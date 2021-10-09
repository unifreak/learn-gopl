package main

import (
	"fmt"
)

// Remove and preserve order.
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Remove and don't need preserve order.
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	fmt.Printf("%q\n", append(runes, runes...))

	fmt.Printf("%q\n", []rune("Hello, 世界"))

	// Use slice as stack
	var stack []int
	stack = append(stack, 1)  // push
	stack = append(stack, 2)
	top := stack[len(stack)-1] // top
	stack = stack[:len(stack)-1] // pop
	fmt.Printf("top=%d, %v\n", top, stack)

	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("%v\n", remove(s, 2))

	s = []int{5, 6, 7, 8, 9}
	fmt.Printf("%v\n", remove2(s, 2))
}