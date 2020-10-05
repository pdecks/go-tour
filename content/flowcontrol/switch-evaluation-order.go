// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

	// note that time in the Go playground (online version of this tour) appears to start at 2009-11-10 23:00:00 UTC,
	// which is when Google officially announced the release of GoLang https://techcrunch.com/2009/11/10/google-go-language/
	today := time.Now().Weekday()

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
