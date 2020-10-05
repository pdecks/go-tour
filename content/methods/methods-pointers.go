// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with type receiver (value)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method with pointer receiver, mutates type instance
// with a value receiver, the Scale method operates on a COPY of the original Vertex value
// the Scale method MUST have a pointer receiver (*Vertex) to change the Vertex
// value declared in the main function
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
