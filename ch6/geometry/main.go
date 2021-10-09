package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	// Calling methods with a pointer receiver
	// 1.
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"
	// 2.
	p = Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"
	// 3.
	p = Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"
	// 4. shortcut.
	// Works when
	// - receiver is variable of type T, receiver parameter has type *T
	// - receiver is variable of type *T, receiver parameter has type T
	// DON'T work if receiver parameter is non-addressable, like a Point literal
	p = Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

}

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
//
// NOTE: to avoid ambiguities, method declarations are not permitted
// on named types that are themselves pointer types.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// define a method with a pointer receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

