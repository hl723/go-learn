package main

import "golang.org/x/tour/pic"
import (
	"image"
	"image/color"
)

type Image struct{
	x, y int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel	
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.x, img.y)	
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}	
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
