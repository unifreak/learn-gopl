package main

import (
	"fmt"
	"time"
)

// define a struct with named type `Employee`
type Employee struct {
	// field order is significant to type identity
	// different order or combination is different type
	//
	// a struct type may contain a mixture of exported and unexported filelds
	//
	// a named struct type S can't declare a field of the same type S: an
	// 	aggregate value cannot contain itself. But S may declare a field
	// 	of pointer type *S, which let us create recursive data structure
	// 	like linked lists and trees
	// 	@see treesort
	ID 		   		 int
	Name, Address    string // consecutive fileds of the same type
						    // may be combined
	Bob 	  		 time.Time
	Position  		 string
	Salary    		 int
	ManagerID 		 int
}

type Point struct{ X, Y int }

func main() {
	var dilbert Employee

	// Access
	// 1. dot notation
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	// 2. through a pointer
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	// 3. a pointer to a struct
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// same as: (*employeeOfTheMonth).Position += ...

	fmt.Printf("%v\n", dilbert)

	// Zero value for a struct is composed of the zero values of each of its fields

	// `empty struct`: has size zero and carries no information
	set := make(map[string]struct{})
	fmt.Printf("%v\n", set)

	// Set value with `struct literals`. it has two forms.
	// 1. with first form, we must specify every vield, in the right order
	//    tends to be used only within the package that defines the struct type
	//    or with smaller struct types for which there is an obvious filed ordering
	//    conventioin, like image.Point{x, y} or color.RGBA{red, green, blue, alpha}
	p := Point{1, 2}
	// 2. with second form, order doesn't matter
	p = Point{X:1}
	// the two forms cannot be mixed in the same literal

	fmt.Println(Scale(p, 5))

	// Shorthand notation to create and init a struct and obtain its address
	pp := &Point{1, 2}
	fmt.Printf("%v\n", pp)

	// If all fiels are comparable, then struct itself is comparable
	// And may be used as the key type of a map
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Printf("%v\n", hits)

	// `Struct embedding`: lets us use one named struct type as
	//    an anonymous field of another struct type, providing
	//    a convenient syntactic shortcut so that a simple dot
	//    expression like x.f can stand for a chain of fileds
	//    like x.d.e.f
	// `Anonymous fileds`: field with a type but no name. The
	//    type of the field must be a named type or a pointer to
	//    a named type
	// @see embed
}

// Pass as arguments and return from function
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// For efficiency, larger struct types are usually passed to or returned
// from functioins using a pointer
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}
// And passing pointer is required if must modify its argument
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
