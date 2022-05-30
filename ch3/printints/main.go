// intsToString is like fmt.Sprintf(values) but adds commas.
package main

import (
	"fmt"
	"bytes"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	// When appending the UTF-8 encoding of an arbitrary rune to a bytes.Buffer,
	// it's best to use bytes.Buffer's WriteRune method, but WriteByte is fine
	// for ASCII characters such as '[' and ']'
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}