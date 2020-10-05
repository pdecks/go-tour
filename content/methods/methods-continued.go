// +build OMIT

package main

import (
	"fmt"
	"math"
)

// you can declare a method on non-struct types, too
type MyFloat float64

// BUT you can ONLY declare a method with a receiver whose type is DEFINED IN THE SAME package
// as the method. therefore, you cannot declare a method with a receiver whose type is defined
// in another package (including built-in types, int, float64 etc) ... which is why this code uses
// the wrapper type MyFloat to create the method taking MyFloat acting on a float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
