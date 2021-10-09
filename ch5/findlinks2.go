package main

import (
	"fmt"
	"net/http"
	"os"
	"log"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		// A multi-valued call may appear as the sole argument when calling
		// a function of multiple parameter. This feature is sometimes convenient
		// during debugging since it lets us print all the results of a calling
		// using a single statement
		//
		// 		log.Println(findLinks(url))
		//
		// is same as
		//
		// 		links, err := findLinks(url)
		// 		log.Println(links, err)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlink2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// The result of a multi-valued call may itself be returned froom a
// (multi-valued) calling function.
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLinks(url)
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		// We must ensure that resp.Body is closed so that network resources
		// are properly released even in case of error.
		// Go's garbage collector recycles unused memory, but do not assume
		// it will release unused operating system resources like open files
		// and network connections. They should be closed explicitly.
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	// HTML parser can usually recover from bad input and construct a
	// document containing error nodes, so Parse rarely failes; when it
	// does, it's typically due to underlying I/O errors.
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}