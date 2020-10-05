// +build OMIT

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// return a map of the counts of each "word" in the string s
func WordCount(s string) map[string]int {
	// initialize the map of char counts
	wc := make(map[string]int)

	for _, c := range strings.Fields(s) {
		if _, ok := wc[c]; ok == true {
			wc[c] += 1
		} else {
			wc[c] = 1
		}
	}

	return wc
}

func main() {
	// wc.Test(_func_) runs a test suite against the provided function and prints success or failure
	wc.Test(WordCount)
}
