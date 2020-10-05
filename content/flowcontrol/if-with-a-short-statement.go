// +build OMIT

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// variables declared by the statement are only in scope until the end of the _if_.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// trying to return v here would fail
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
