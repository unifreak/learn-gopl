// Ex3.10 inserts commas in a non-negative decimal integer string
// in non-recursive way.
package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Println(comma(arg))
    }
}

func comma(s string) string {
    var buf bytes.Buffer
    lead := len(s) % 3
    if lead == 0 {
        lead = 3
    }
    buf.WriteString(s[:lead])

    for i := lead; i < len(s); i += 3 {
        buf.WriteString(", ")
        buf.WriteString(s[i:i+3])
    }

    return buf.String()
}