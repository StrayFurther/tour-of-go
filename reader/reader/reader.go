package main

import "golang.org/x/tour/reader"
import "io"

type MyReader struct{}

func (m *MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), io.EOF
}
func main() {
	reader.Validate(&MyReader{})
}
