// +build no-build OMIT

package main

import (
	"fmt"
	"math"
)

// an interface type is defined as a set of METHOD signatures
// a value of interface type can hold any value that implements those methods
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// results in error: "cannot use v (type Vertex) as type Abser in assignment:
	//   Vertex does not implement Abser (abs method has pointer receiver)
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloat implements the Abs() method defined in the Abser interface above
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
