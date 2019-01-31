package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot* rot13Reader) Read(x []byte) (int, error) {
	m := make([]byte, len(x))
	s, _ := rot.r.Read(m)
	for k, v := range m {
		// fmt.Print(v)
		switch {
			case v > 96 && v < 110:
				x[k] = v + 13
			case v < 123 && v >= 110:
				x[k] = v - 13
			case v > 65 && v < 79:
				x[k] = v + 13
			case v < 90 && v >= 79:
				x[k] = v - 13
			default:
				x[k] = v
		}
	}

	return s, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
