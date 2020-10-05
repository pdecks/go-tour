// +build OMIT

package main

import (
	"fmt"
	"image"
)

// Package image defines the _Image_ interface
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle // image.Rectangle
//     At(x, y int) color.Color
// }
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
