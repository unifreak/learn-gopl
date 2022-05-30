package pkg

import "fmt"

type point struct {
	X, Y int
}

type Circle struct {
	point
	Radius int
}

func init() {
	var c Circle
	c.X = 1
	fmt.Printf("from pkg: %#v", c)
}
