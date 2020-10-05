// +build OMIT

package main


import "fmt"

type Person struct {
	Name string
	Age  int
}

// one of the most ubiquitous interfaces is _Stringer_, defined by the fmt package
// a _Stringer_ is a type that can describe itself as a string. The fmt package (and others)
// look for this interface to print values ...
// type Stringer interface {
//	String() string
//}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
