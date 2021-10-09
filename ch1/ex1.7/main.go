// Ex1.7: use io.Copy(dst, src) to avoid requiring a buffer large enough to hold the entire stream

// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1.7: %v\n", err)
			os.Exit(1)
		}
		// way to check `io.Copy()` error
		//
		// if use _, err := io.Copy()...
		// will throw error: no new variables on left side of :=
		// because of both _ and err are NOT new variables
		// so need to use = instead :=
		//
		// see https://stackoverflow.com/questions/41574028/golang-no-new-variables-on-left-side-of-while-a-similar-one-didnt-occur-th/41574085
		//
		// BUT @? why `if _, err := io.Copy(...); err != nil {}` can get away?
		// probably becuase inside if block, _ and err are new variables?
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1.7: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}