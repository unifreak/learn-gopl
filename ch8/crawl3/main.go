// crawl3 solve the problem of excessive concurrency with 20 long-lived crawler. but
// it doesn't address the problem of program termination.
//
// goroutines. this version uses the original crawl function that has no counting
// semaphore, but calls it  from one of 20 long-lived crawler goroutines, thus
// ensuring that at most 20 HTTP requests are active concurrently.
package main

import (
	"os"
	"fmt"
	"log"

	"learn/gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string) // list of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-lin arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				worklist <- foundLinks
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}