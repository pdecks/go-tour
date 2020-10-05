// +build OMIT

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP ... EXCEPT that Go ONLY runs the SELECTED case,
	// NOT all the cases that follow. In essence, the _break_ statement provided at the end of each case in the
	// aforementioned languages is provided automatically in Go
	// NOTE: Go's switch cases need not be constants, and the values need not be integers
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
