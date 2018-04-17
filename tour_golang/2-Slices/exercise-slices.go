package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for i := range p {
		for j := range p[i] {
			p[i][j] = uint8(0) + uint8(i)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
