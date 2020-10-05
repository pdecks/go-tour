// +build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// if the top-level type is a type name, it can be omitted from the elements of the literal ... not seeing what that means
// from this example though ...
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
