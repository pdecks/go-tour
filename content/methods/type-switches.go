// +build OMIT

package main

import "fmt"

type MyFloat float64

func do(i interface{}) {
	// type switch: permits several type assertions in series
	// ~ regular switch, except the cases specify TYPES, not values.
	//   the type (value) is compared against the type held by the given interface value
	// the declaration in a type switch has the same syntax as a type assertion, i.(T), except
	// the specific type T is replaced with the keyword _type_ -> i.(type)
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case MyFloat:
		fmt.Printf("Hello, %v is a %T", v, v)
		// in the default case, the variable v is of the same interface type and value as i
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		describe(i)
	}
}

func describe(i interface{}) {
	fmt.Printf("Interface tuple: (%v, %T)\n", i, i)
}

func main() {
	do(21)
	do("hello")
	do(true)
	do( MyFloat(3.14))
}
