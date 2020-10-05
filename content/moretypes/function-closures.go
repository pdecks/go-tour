// +build OMIT

package main

import "fmt"

// Go functions may be closures, which are function values that reference variables from outside their bodies. The closure
// may access and assign to the referenced variables; in this sense, the function is "bound" to the variables

// returns a closure, where each closure is bound to its own _sum_ variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
