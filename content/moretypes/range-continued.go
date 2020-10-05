// +build OMIT

package main

import "fmt"

// you can skip the index or value in a range iterator by using _ (anonymous)
// for i, _ := range pow -> if you only want the index, you can omit the element variable -> for i := range pow
// for _, value := range pow
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
