// +build OMIT

package main

import "fmt"

// _make_ is used to create dynamically sized arrays
// allocates a zero'd array and returns a slice referencing that array
// make(type, len, cap)
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	//b = b[:cap(b)]
	//printSlice("b", b)
	//b = b[1:]
	//printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
