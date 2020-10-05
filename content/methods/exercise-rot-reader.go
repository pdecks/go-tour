// +build no-build OMIT

package main

import (
	"io"
	"os"
	"strings"
)

// A common pattern is an io.Reader wrapping another io.Reader to modify the stream
// e.g. gzip.NewReader -> takes io.Reader (data stream) and returns a stream of the decompressed data
type rot13Reader struct {
	r io.Reader
}

// implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream
// by applying the rot13 substitution cipher to all alphabetical characters
func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i]) // apply cipher
	}
	return
}

// rot13 substitution cipher
func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b<= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b - a + 13) % (z - a + 1) + a
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
