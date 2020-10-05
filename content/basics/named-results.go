// +build OMIT

package main

import "fmt"

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = 1 // modification
	return
}

func main() {
	fmt.Println(split(17))
}
