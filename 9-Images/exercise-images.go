package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Rect image.Rectangle
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.Rect
}

func (Image) At(x, y int) color.Color {
	v := uint8(x)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{image.Rect(0, 0, 256, 256)}
	pic.ShowImage(m)
}
