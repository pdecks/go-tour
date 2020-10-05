// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we DO NOT need to explicitly declare that it does so (i.e. do not use "implements" keyword)
// this decouples the definition of an interface from its implementation which could
// then appear in any package without prearrangement
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
