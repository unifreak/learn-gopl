// The funciton call io.Copy(dst, src) read from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without
// requireing a buffer large enough to hold the entire stream. Be sure to check
// the error result of io.Copy.
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
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1.7: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}