// Enhance comma so that it deals correctly with floating-point numbers and an
// optional sign
package main

import (
    "bytes"
    "fmt"
    "os"
    "strings"
)

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Println(comma(arg))
    }
}

func comma(s string) string {
    var buf bytes.Buffer
    if s[0] == '-' || s[0] == '+' {
        buf.WriteByte(s[0])
        s = s[1:]
    }

    lead := len(s) % 3
    if lead == 0 {
        lead = 3
    }
    buf.WriteString(s[:lead])

    end := strings.Index(s, ".")
    if end == -1 {
        end = len(s)
    }

    for i := lead; i < end; i += 3 {
        buf.WriteString(", ")
        buf.WriteString(s[i:i+3])
    }
    buf.WriteString(s[end:])

    return buf.String()
}