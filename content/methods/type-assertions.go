// +build no-run OMIT

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// a _type assertion_ provides access to an interface value's underlying concrete value
	// the below statement asserts that the interface value i holds the concrete type _string_ and
	// assigns the underlying _string_ value to the variable s
	s := i.(string)
	fmt.Println(s) // "hello"

	// a type assertion can return two values:
	// 1. the underlying value
	// 2. a boolean reporting whether the assertion succeeded
	// NOTE: similarity of this syntax and reading from a map
	// below, if i holds a _string_, then _s_ will be the underlying value and _ok_ will be _true_
	s, ok := i.(string)
	fmt.Println(s, ok)

	// ... if not, _ok_ will be _false_ and _f_ will be the zero value of type _float64_ and no panic occurs
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// if i does not hold a type _T_, the statement will trigger a panic
	f = i.(float64) // panic ... no panic if you comment out only this line
	fmt.Println(f)
}
