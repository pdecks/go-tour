// +build OMIT

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// variables declared inside an _if_ short statement are also available inside any of the _else_ blocks
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	// both calls to _pow_ return their results before the call to fmt.Println begins
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
