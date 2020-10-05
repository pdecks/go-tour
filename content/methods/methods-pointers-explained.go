// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// removing the pointer receiver (*Vertex) and the pointer argument to the call to
// Scale() in main (&v), this function is now operating on a COPY of v, returning the
// abs(v) ->5 instead of 50, using the pointer receiver and argument.
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
