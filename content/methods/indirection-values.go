// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	p := &Vertex{4, 3}
	// pointer argument passed to value receiver
	// METHODS with value receivers may take either a value or a pointer as the receiver
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p)) // *p, the variable to which p points ... a Vertex value


	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // value receiver

	// functions that take a value argument MUST take a value of that type
	// fmt.Println(AbsFunc(&v)) -> compile error!
	fmt.Println(AbsFunc(v))


}
