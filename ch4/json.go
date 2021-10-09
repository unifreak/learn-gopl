package main

import (
	"fmt"
	"encoding/json"
	"log"
)

/**
 * JSON: JavaScript Object Notation
 * - basic json type:
 * 		numbers (decimal / scientific)
 * 		booleans (true / false)
 * 		string. unlike Go, \Uhhhh denote UTF-16 codes, not runes.
 * - array: ordered, to encode Go arrays and slices.
 * - object: to encode Go maps (with string keys) and structs.
 */

type Movie struct {
	Title 	string
	// `json:"released"` is a `field tag`
	//
	// field tag: a string of metadata associated at compile time with
	// 		the field of a struct. it can be any literal string, but
	// 		conventionally interpreted as space-separated list of `key:"value"`
	// 		pairs; since they contain double quotation marks, field tags
	// 		are usually written with raw string literals.
	//
	// "json" key control the behavior of the `encoding/json` package
	// "released" specify alternative JSON name for Go field
	// "omitempty" indicate no JSON output if the field has zero value
	Year  	int 	`json:"released"`
	Color 	bool 	`json:"color,omitempty"`
	Actors  []string
}

var movies = []Movie{
	{Title: "Casablance", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	// Converting Go structure to JSON is callled `marshaling`, done by `json.Marshal`
	// only exprted fields are marshaled
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// For human consumption, use `json.MarshalIndetn`
	// "": line prefix
	// "    ": indentation
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// The inversion operation to marshaling is called `Unmarshaling`
	// done by `json.Unmarshal`.
	//
	// We select which part to decode by defining suitable Go data structures.
	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarhsaling failed: %s", err)
	}
	fmt.Println(titles)
}