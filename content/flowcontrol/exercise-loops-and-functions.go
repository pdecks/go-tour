// +build no-build OMIT

package main

import (
	"fmt"
	"math"
)

// implement a square root function
// given a number x, find the number z for which z^2 is most nearly x
// typically, this is computed using a loop. starting with a guess, z, we can adjust z based on how close z^2 is to x
func Sqrt(x float64) float64 {
	z := 1.0 // first guess
	err := float64((z*z - x) / (2.0*z))
	// to begin, repeat the calculation 10 times, printing each z to observe how accurate the approximation is
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("Start of loop for iteration %v: \n", i)
	//	fmt.Printf("err: %v\n", err)
	//	z -= err
	//	fmt.Printf("Updated z: %v\n", z)
	//	err = (z*z - x) / (2.0*z)
	//}

	for math.Abs(err) > 0.005 {
		fmt.Printf("err: %v\n", err)
		z -= err
		fmt.Printf("z: %v\n", z)
		err = (z*z - x) / (2.0*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
