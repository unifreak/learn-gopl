package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create map.
	// 1. with make()
	ages := make(map[string]int) // maping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34

	// 2. with `map literal`
	ages = map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	var empty = map[string]int{}

	fmt.Printf("%v\n%v", ages, empty)

	// Access
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // 32

	// Remove
	delete(ages, "alice")

	// All above ops is safe even element isn't in the map.
	// The value will be the zero value for its type.
	ages["bob"] = ages["bob"] + 1
	// Sometime you may have to distinguish between a nonexist element
	// and an element that happens to have the zero value
	if _, ok := ages["david"]; !ok {
		// handle nonexist element
	}

	// shorthand assignment also works
	ages["bob"] += 1
	ages["bob"]++
	// but element is not a varialbe, we cannot take its address
	// (growing a map might cause rehashing, invalidating the address)

	// Enumerate
	//   order of map iteration is unspecified, in practice, is random
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	//   to make key/value in order, we must sort the keys explicitly
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	//   since we know the size from the outset, it's more efficient to
	//   allocate an array of the required size up front.
	names = make([]string, 0, len(ages))

	// Zero value for map type is nil. a reference to no hash table at all.
	var neo map[string]int
	fmt.Println(neo == nil) // "true"
	fmt.Println(len(neo) == 0) // "true"
	// its safe to perform lookup, delete, len, and range on nil map
	// since it behave like an empty map
	// But storing to a nil map cause a panic
	// you must allocate the map before storing into it
	// neo["carol"] = 21 // panic: assignment to entry in nil map

	// Go have no `set`, but map can serve as `set`
	// see dedup.go
}

// maps cannot be compared, except with `nil`
// we must write a loop to do this:
func equal(x, y map[string]int) bool {
	if (len(x) != len(y)) {
		return false
	}
	for k, xv := range x {
		// Note how we use !ok to distinguish nonexist and zero value element
		// if writen as
		// 		xv != y[k]
		// then
		// 		equal(map[string]int{"A": 0}, map[string]int{"B": 42})
		// will report as equal, incorrectly
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// Since map's keys must be comparable, if we need a map or set
// whose keys are slices, we need to do it like this:
var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string) { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

// The value type of a map can itself be a composite type
// such as a map or slice
// see graph.go