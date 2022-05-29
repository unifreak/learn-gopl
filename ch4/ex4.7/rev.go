// Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?
package rev

import (
    "unicode/utf8"
)

func rev(s []byte) {
    size := len(s)
    for i := 0; i < size/2; i++ {
        s[i], s[size-i-1] = s[size-i-1], s[i]
    }
}

func reverse(s []byte) []byte {
    for i, _ := range s {
        _, size := utf8.DecodeRune(s[i:])
        rev(s[i:i+size])
    }
    rev(s)
    return s
}