// Compare loop and signle expression performance of pop count.
package main

import (
	"fmt"
	"./popcount"
	"time"
)

const (
)

func main() {
	var runs uint64 = 10000000
	var x uint64

	start := time.Now()
	for x = 0; x < runs; x++ {
		popcount.PopCount(x)
	}
	fmt.Printf("single expression pop count take %f seconds to run\n", time.Since(start).Seconds())

	start = time.Now()
	for x = 0; x < runs; x++ {
		popcount.LoopPopCount(x)
	}
	fmt.Printf("loop expression pop count take %f seconds to run\n", time.Since(start).Seconds())

	start = time.Now()
	for x = 0; x < runs; x++ {
		popcount.ShiftPopCount(x)
	}
	fmt.Printf("shift expression pop count take %f seconds to run\n", time.Since(start).Seconds())
}
