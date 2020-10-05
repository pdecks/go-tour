// +build no-build OMIT

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // exported names start with a capital letter, so this should be "Pi" not "pi"
}
