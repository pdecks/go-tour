// +build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// maps keys -> values
// the zero value of a map is _nil_. a _nil_ map has NO keys nor can keys be added
var m map[string]Vertex

func main() {
	// make returns a map of the given type, initialized and ready for use
	// string -> Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
