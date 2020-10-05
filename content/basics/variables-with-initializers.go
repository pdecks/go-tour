// +build OMIT

package main

import "fmt"

var i, j int = 1, 2 // if an initializer is present, the type can be omitted. therefore, this can be 'var i, j = 1, 2'

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
