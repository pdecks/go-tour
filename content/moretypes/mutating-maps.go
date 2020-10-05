// +build OMIT

package main

import "fmt"

func main() {
	m := make(map[string]int)

	// insert of update: m[key] = elem
	// retrieve: elem = m[key]
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete an element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// test that a key is present with two-value assignment
	// key present ? ok == true : ok == false
	// if key not present, elem is the zero value for the map's element type
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
