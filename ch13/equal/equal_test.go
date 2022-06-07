package equal

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))        // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))      // "false"
	fmt.Println(Equal([]string(nil), []string{}))             // "true"
	fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"

	// Circular linked lists a -> b -> a and c -> c
	type link struct {
		value string
		tail *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) 		// "true"
	fmt.Println(Equal(b, b)) 		// "true"
	fmt.Println(Equal(c, c)) 		// "true"
	fmt.Println(Equal(a, b)) 		// "false"
	fmt.Println(Equal(a, c)) 		// "false"
}
