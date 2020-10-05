// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("v1: %v\n", v1)
	q := &v1
	q.X = 12
	fmt.Printf("v1 updated through pointer: %v\n", v1)
	fmt.Println(v1, p, v2, v3) // showing that p is pointing to a different Vertex instance ...
}
