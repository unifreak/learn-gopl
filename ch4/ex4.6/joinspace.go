// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
package joinspace

import (
    "unicode"
    "unicode/utf8"
)

func joinspace(s []byte) []byte {
    n := 0
    seen := false
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRune(s[i:])
        if (unicode.IsSpace(r)) {
            if seen {
                i += size
                continue
            }
            seen = true
            s[n] = ' '
            n += 1
        } else {
            seen = false
            copy(s[n:], s[i:])
            n += size
        }
        i += size
    }
    return s[:n]
}