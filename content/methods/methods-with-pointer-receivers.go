// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// there are two reasons to use a pointer receiver:
// 1. so the method can modify the value its receiver points to
// 2. to avoid copying the value on each method call. this can be more efficient if the
//    receiver is a large struct, for example
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// even though Abs doesn't need to modify its receiver, in generall ALL methods on a given
// type should have either value or pointer receivers, but not a mixture of both
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
