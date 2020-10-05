// +build OMIT

package main

import "fmt"

// the type []T is a slice with elements of type T
// specifies two indices, a low and a high bound, high excluded
// a slice is a dynamically sized flexible view into the elements of an array
// a slice does not store any data -- it is a view
// changing the elements of a slice modifies the corresponding elements of its underlying array
// other slices that share the same underlying array will see those changes
//
// defaults: low - zero; high: length of array
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // elements 1 through 3 of a, where the array is zero-based
	fmt.Println(s)
}
