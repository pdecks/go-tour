// +build OMIT

package main

import "fmt"

func main() {
	// the interface type that specifies zero methods is known as the _empty interface_
	// interface{} ... an empty interface my hold values of any type, as every type implements
	// at least zero methods
	// empty interfaces are used by code that handles values of unknown type, e.g. fmt.Print
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
	// from fmt.Printf ...
	// func Printf(format string, a ...interface{}) (n int, err error) {
	//	return Fprintf(os.Stdout, format, a...)
	//}
}
