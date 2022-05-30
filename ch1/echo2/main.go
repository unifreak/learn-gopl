// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    // range produces a pair of values: index and value
    // go deso not permit unused local vaiables, so use _ the blank identifier
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

// if data is large, use + will be costly, see echo3.go