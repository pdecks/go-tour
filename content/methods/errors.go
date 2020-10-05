// +build OMIT

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// Go programs express error state with _error_ values. The _error_ type is a built-in interface
// similar to fmt.Stringer. The fmt package looks for the _error_ interface when printing values
// type error interface {
//     Error() string
//}
func main() {
	// functions often return an error value, and calling code should handle errors by testing
	// whether the error equals nil
	// a nil _error_ denotes success; a non-nil _error_ denotes failure
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
