package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(a []byte) (int, error) {
	n, err := r.r.Read(a)

	for i, c := range a {
		switch {
		case 'A' <= c && c <= 'M', 'a' <= c && c <= 'm':
			a[i] += 13
		case 'N' <= c && c <= 'Z', 'n' <= c && c <= 'z':
			a[i] -= 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
