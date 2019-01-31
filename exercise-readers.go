package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(x [] byte) (int, error) {
	for i, _ := range x {
		x[i] = 'A'
	}
	
	return len(x), nil
}

func main() {
	reader.Validate(MyReader{})
}
