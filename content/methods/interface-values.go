// +build OMIT

package main

import (
	"fmt"
	"math"
)

// under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// wrapper type because the method must be defined in the same package as the type definition
// meaning we cannot write methods that take built-in types (e.g. float, int)
type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	// calling a method on an interface value executes the method of the same name
	// on its UNDERLYING TYPE -> i.e. the line below executes (t *T) M()
	// kind of like "overloading" methods
	i.M() // interface value (&T{"Hello"}, *T)

	i = F(math.Pi)
	describe(i)
	// ... while the line below executes (f F) M()
	i.M() // interface value (F(math.Pi), F)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
