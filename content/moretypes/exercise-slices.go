// +build no-build OMIT

package main

import "golang.org/x/tour/pic"

// implement - return a slice of length _dy_, where each element is a slice of _dx_ 8-bit unsigned int
//
func Pic(dx, dy int) [][]uint8 {
	// use uint8(_some_int_value_) to convert from int -> uint8

	// allocate a two-dimensional slice of length dy
	p := make([][]uint8, dy)

	// allocate a slice for each row
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}

	// drawing
	for y := range p {
		for x := range p[y] {
			// p[y][x] = uint8(x^y);
			p[y][x] = uint8(x*y) // ooooo so pretty!
			//p[y][x] = uint8((x+y)/2) // boooooring gradient
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
