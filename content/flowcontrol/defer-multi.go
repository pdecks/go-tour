// +build OMIT

package main

import "fmt"

// see also: https://blog.golang.org/defer-panic-and-recover
// defer is commonly used to simplify functions that perform various clean-up actions
// ex in the blog post -> copying one file to a new file, using defer to Close() the files
// ... kind of like _finally_ block
// defer is also commonly used to release a mutex, print a footer, etc.
//
// deferred functions may read and assign to the returning function's named return value, which can be convenient for
// modifying the error return value of a function
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println("counting")

	// deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in
	// last-in-first-out order
	//

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // a deferred function's arguments are evaluated when the defer statement is evaluated
	}

	fmt.Println("done ... but deferred print statements execute when main() finishes")

	fmt.Printf("Call to c(): %v\n", c())
}
