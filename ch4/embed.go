package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8 // equivalent to w.Circle.Point.X = 8
	w.Y = 8 // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	// explicit forms shown in the comment still valid.
	// hence, anonymous field do have names (the named type)
	// but those names are optional in dot expression
	// so you can't have two anonymous field of the same type.
	// also the visibility of the field is determined by its type name.

	// But there is no corresponding shorthand for the struct literal
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here
	}
	fmt.Printf("%#v\n", w) // # adverb cause %v to display values in
						   // a form similar to Go syntax

	w.X = 42
	fmt.Printf("%#v\n", w)
}