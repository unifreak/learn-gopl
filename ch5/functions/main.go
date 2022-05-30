package main

import (
	"fmt"
	"math"
)

/**
 * function declaration
 *
 * 		func name(parameter-list) (result-list) {
 * 			body
 * 		}
 *
 * funcitons may be recursive. see findlinks1.go and outline.go
 */

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int { return x + y }

// same type can be factored so type itself is written only once.
// results may be named, each name declares a local variable initialized to zero value.
func sub(x, y int) (z int) { z = x-y; return }

// _ can be used to emphasize that a parameter is unused.
func first(x int, _ int) int { return x }

func zero(int, int) int { return 0 }

// function declaratioin without a body, indicating that the function is implemented in
// a language other than Go.
func Sin(x float64) float64 { }

// In a function with named results, the oprands of a return statement may
// be omitted. This is called bare return. Bare return can reduce code cuplication,
// but they rarely make code easier to understand. Best use them sparingly.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return // 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return // 0, 0, err
	}
	words, images = CountWordsAndImages(doc)
	return // words, images, err
}

func CountWordsAndImages(n *html.Node) (word, image int) {
	/* ... */
}

func main() {
	// Go has no concept of default parameter values, nor any way to specify arguments by name.
	// arguments are passed by value.
	fmt.Println(hypot(3, 4)) // 5

	// type of funciton is sometimes called its `signature`.
	// parameter and result's name don't affect the type, noe does whether or not
	// they  are declared using the factored form.
	fmt.Printf("%T\n", add) 	// "func(int, int) int"
	fmt.Printf("%T\n", sub) 	// "func(int, int) int"
	fmt.Printf("%T\n", first) 	// "func(int, int) int"
	fmt.Printf("%T\n", zero) 	// "func(int, int) int"
}
