package main

import (
	"fmt"
	"io/ioutil"
	memo "learn/gopl/ch9/memo1"
	"log"
	"net/http"
	"time"
)

func main() {
	m := memo.New(httpGetBody)
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func httpGetBoy(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
