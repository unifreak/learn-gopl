// The du2 command computes the disk usage of the files in a directory.
//
// It prints the totals periodically, but only if the -v flag is specified.
//
// The background gorouine that loops over roots remains unchanged. The main goroutin
// now uses a ticker to generate events every 500ms, and a select statement to wait
// for either a file size message, in which case it updates the totals, or a tick
// event, in which case it prints the current totals.
//
// If the -v flag is not specified, the tick channel remains nil, and its case in the
// select is effectively disabled.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"path/filepath"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the inital directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		// since the program no longer uses a range loop, the first select case must
		// explicitly test whether the fileSizes channel has been closed, using ok.
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}
