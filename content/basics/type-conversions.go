// +build OMIT

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var a string = string(x) // doesn't work
	fmt.Println(x, y, z, a)
}
