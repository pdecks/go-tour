// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Printf("read i through pointer: %v\n", *p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Printf("set i through pointer: %v\n", i)  // see the new value of i

	p = &j         // point to j
	*p /= 37       // divide j through the pointer
	fmt.Printf("point p to j and update j: %v", j) // see the new value of j
}
