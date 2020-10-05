// +build OMIT

package main

import "fmt"

func main() {
	sum := 1
	// omitting the init and post statements essentially converts _for_ to _while_
	// .. and you can drop the semicolons (for-is-gos-while.go)
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
