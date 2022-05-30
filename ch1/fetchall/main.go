// Fetchall fetches URLs in parallel and reports their times and sizes
//
// gorutine: a concurrent function execution
//
// channel: communication mechanism that allows one goroutine to pass values of
// a specified type to another goroutines
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() { // main() run in goroutine
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // go create additional goroutines
	}
	// When one goroutine attemps a send or receive on an channel, it blocks
	// until another goroutine attempts the corresponding recive or send
	// operation. Hence having main() do all printing ensures that output from
	// each goroutine is processed as a unit, with no danger of interleaving if
	// two goroutines finish at the same time
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}