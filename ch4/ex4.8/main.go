// Modify charcount to count letters, digits, and so on in their Unicode categories,
//  using functions like unicode.IsLetter.
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

var (
    counts = make(map[rune]int)             // counts of Unicode characters
    countsPerType = make(map[string]int)    // countes for category
    utflen [utf8.UTFMax+1]int               // count of lengths of UTF-8 encodings. UTFMax is 4
    invalid = 0                             // count of invalid UTF-8 characters
)

func countType(r rune) {
    if unicode.IsControl(r) {
        countsPerType["Control"]++
    }
    if unicode.IsDigit(r) {
        countsPerType["Digit"]++
    }
    if unicode.IsGraphic(r) {
        countsPerType["Graphic"]++
    }
    if unicode.IsLetter(r) {
        countsPerType["Letter"]++
    }
    if unicode.IsLower(r) {
        countsPerType["Lower"]++
    }
    if unicode.IsMark(r) {
        countsPerType["Mark"]++
    }
    if unicode.IsNumber(r) {
        countsPerType["Number"]++
    }
    if unicode.IsPrint(r) {
        countsPerType["Print"]++
    }
    if unicode.IsPunct(r) {
        countsPerType["Punct"]++
    }
    if unicode.IsSpace(r) {
        countsPerType["Space"]++
    }
    if unicode.IsSymbol(r) {
        countsPerType["Symbol"]++
    }
    if unicode.IsTitle(r) {
        countsPerType["Title"]++
    }
    if unicode.IsUpper(r) {
        countsPerType["Upper"]++
    }
}

func main() {
    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()  // return rune, nbytes, error
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 { // not a legal UTF-8 encoding
            invalid++
            continue
        }
        counts[r]++
        countType(r)
        utflen[n]++
    }
    fmt.Printf("rune\tcount\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Print("\nlen\tcount\n")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    fmt.Print("\ncategory\tcount\n")
    for t, n := range countsPerType {
        fmt.Printf("%q\t%d\n", t, n)
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}