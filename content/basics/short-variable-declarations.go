// +build OMIT

package main

import "fmt"

// inside a function, the short assignment statement (:=) can be used in place of a _var_ declaration with implicit type
// outside a function, every statement begins with a keyword (_var_, _func_, etc.) so the := construct is NOT available
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
