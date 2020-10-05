// +build OMIT

package main

import "fmt"

func deferPlay() {
	defer fmt.Println("this is another deferral") // returns second
	fmt.Println("check this out") // returns first
}
func main() {
	// _defer_ feders the execution of a function until the surrounding function returns
	// the deferred call's arguments are evaluated immediately, but the function call is not executed until the
	// surrounding function returns
	defer fmt.Println("world") // won't execute until main() returns, returns last

	// expect: "check this out\n this is another deferral\n hello\n person from another\n world"
	deferPlay() // returns first because the deferral inside is wrapped only by deferPlay(), not main()


	fmt.Println("hello") // returns third
	fmt.Println("person from another") // returns fourth
}
