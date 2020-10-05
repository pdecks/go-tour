// +build OMIT

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// passing e directly will send the program into an infinite loop ... ?
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
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
	return z, nil
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) // prints error message
}
