// Ex1.4 print the names of all files in which each duplicaated line occurs.
package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    counts = make(map[string]int)
    seenIn = make(map[string]string)
)

func main() {
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin)
    } else {
        for _, arg := range files {
            // .Open() return two value: `*os.File` and built-in `error` type
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
                continue
            }
            countLines(f)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
            fmt.Printf("seen in: %s\n", seenIn[line])
        }
    }
}

// `func` can be declared afterwards
func countLines(f *os.File) {
    // `map` is a reference to data structure create by `make()`
    // hence change ot it is visible to callers
    input := bufio.NewScanner(f)

    var seen = make(map[string]bool)
    for input.Scan() {
        line := input.Text()
        counts[line]++
        if !seen[line] {
            seenIn[line] += " " + f.Name()
            seen[line] = true
        }
    }
    // NOTE: ignoring potential errors from input.Err()
}