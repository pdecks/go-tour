// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	// METHODS with a pointer RECEIVER take either a value or a pointer as the receiver
	// as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5), b/c Scale has a
	// pointer receiver
	v.Scale(2)

	// functions with a pointer argument MUST take a pointer
	// ScaleFunc(v, 10) -> compile error!
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
