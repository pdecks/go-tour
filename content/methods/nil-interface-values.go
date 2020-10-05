// +build no-run OMIT

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// a nil interface value holds neither value NOR concrete type
	// calling a method on a nil interface is a runtime error because there is no type
	// inside the interface tuple to indicate which concrete method to call!
	i.M() // "panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
