// Modify the echo program to also print os.Args[0], the name of
// the command that invoked it.
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // to include command name, both os.Args or os.Args[0:] work
    fmt.Println(strings.Join(os.Args, " "))
}