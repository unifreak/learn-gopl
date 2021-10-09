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