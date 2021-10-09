// Ex4.9: Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before the
// first call to Scan to break the input into words instead of lines.
package main

import (
    "sort"
    "io"
    "os"
    "fmt"
    "bufio"
)

var out io.Writer = os.Stdout
var counts = make(map[string]int)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "usage: os.Args[0] file[...]")
    }
    for _, arg := range os.Args[1:] {
        file, err := os.Open(arg)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ex4.9: %v\n", err)
            os.Exit(1)
        }
        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            word := scanner.Text()
            counts[word]++
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintf(os.Stderr, "ex4.9: %v\n", err)
            os.Exit(1)
        }
    }

    var keys []string
    for word, _ := range(counts) {
        keys = append(keys, word)
    }
    sort.Strings(keys)
    for _, word := range(keys) {
        fmt.Fprintf(out, "%s: %d\n", word, counts[word])
    }
}