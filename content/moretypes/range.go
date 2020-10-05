// +build OMIT

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// the _range_ form of a _for_ loop iterates over a slice or map (~ iterator?)
// when ranging over a slice, two values are returned for each iteration - the index, and a COPY of the index's element
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
