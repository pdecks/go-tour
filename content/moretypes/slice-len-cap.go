// +build OMIT

package main

import "fmt"

// a slice has both a _length_ and a _capacity
// length - the number of elements it actually contains -> len(s)
// capacity - number of elements in the underlying array (counting from first element in the slice) -> cap(s)
// you can extend a slice's length by re-slicing it, provided it has sufficient capacity

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	// This reduces the capacity since we start counting from the first element in the underlying array
	// s = s[2:7] // slice bounds out of range [:7] with capacity 6
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
