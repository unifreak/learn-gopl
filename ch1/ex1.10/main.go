// Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes
// much. Do you get the same content each time? Modify fetchall to print its
// output to a file so it can be examined.
package main

import (
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	temp, err := os.CreateTemp("/tmp", urlpkg.PathEscape(url)+"*")
	if err != nil {
		ch <- fmt.Sprintf("while creating file for %s: %v\n", url, err)
	}
	defer temp.Close()
	nbytes, err := io.Copy(temp, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s \n\tsaved in (%s)", secs, nbytes, url, temp.Name())
}
