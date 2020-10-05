// +build OMIT

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	a, b = swap(a, b) // re-assignment with 2nd swap, hence change in notation
	fmt.Println(a, b)
}
