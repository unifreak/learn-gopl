// prints its command-line arguments' index and value one per line
package main

import (
    "fmt"
    "os"
)

func main() {
	for i, cmd := range os.Args[1:] {
		fmt.Println(i, cmd)
	}
}