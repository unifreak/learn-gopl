// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// behind the secenes, the server runs the handler for each incoming request in
	//   a seperate goroutin so that it an serve multiple requests simultaneously.
	// hence we must ensure that at most one goroutine access `count` at a time, to
	//   avoid race condition. That's why we use `mu.Lock` and `mu.Unlock`
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

// counter echos the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}