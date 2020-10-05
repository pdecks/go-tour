// +build OMIT

package main

import (
	"fmt"
	"io"
	"strings"
)

// the io pkg specifies the io.Reader interface, which represents the read end of a data stream
// (the Go std lib contains many implementations of these interfaces, e.g. files, network conn,
//  compressions, ciphers, etc.)
// io.Reader interface has a Read method:
//   func (T) Read(b []byte) (n int, err error)
// Read() populates the given byte slice with data and returns the number of bytes populated and an
// error value. It returns io.EOF when the stream ends

func main() {
	r := strings.NewReader("Hello, Reader!") // starts at index 0, "prevRune" = -1

	b := make([]byte, 8) // 8-byte slice
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // need b[:n] instead of b[:] for the case where n == 0 (EOF)
		if err == io.EOF {
			break
		}
	}
}
