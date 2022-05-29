// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
    "flag"
    "fmt"
    "os"
    "crypto/sha256"
    "crypto/sha512"
)

var flag256 = flag.Bool("256", true, "output sha256")
var flag384 = flag.Bool("384", false, "output sha384")
var flag512 = flag.Bool("512", false, "output sha512")

func main() {
    flag.Parse()

    for _, arg := range flag.Args() {
        if *flag512 {
            fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512([]byte(arg)))
        } else if *flag384 {
            fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum384([]byte(arg)))
        } else {
            fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256([]byte(arg)))
        }
    }
    os.Exit(0)
}