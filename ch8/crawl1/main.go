package main

import (
	"fmt"
	"log"
	"os"

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

// The main function resembles breadthFirst (ch5/findlinks). But this time, instead of
// representing the queue using a slice, we use a channel.
//
// crawl1 has two problem:
//
// 1. it created so many network connections at once that it
// 	  exceeded the per-process limit on the number of open files, causing operations
// 	  such as DNS loopups and calls to net.Dail to start failing.
//
//    it's too parallel. Unbounded parallelism is rarely a good idea since there is always
// 	  a limiting factor in the system.
//
// 2. the program never terminates, even when it has discovered all the links
// 	  reachable from the initial URLs. Wee need to break out of the main loop when the
// 	  worklist is empty and no crawl goroutines are active.
//
// See crawl2, crawl3 for workarounds.
func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	//
	// Must run in tis own goroutine to avoid deadlock, as stuck situation in which
	// both the main goroutine and a crawler goroutine attempt to send to each other
	// while neither is receiving.
	//
	// An alternative solution would be to use a buffered channel.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
