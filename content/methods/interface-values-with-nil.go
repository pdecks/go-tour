// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	// if the concrete value inside the interface itself is nil, the method will be called
	// with a nil receiver
	// interface value (*T<nil struct>, *main.T) [in package main]
	// NOTE: an interface value that holds a nil concrete value (*T) is ITSELF NON-NIL
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M() // interface value (&{hello}, *main.T)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
