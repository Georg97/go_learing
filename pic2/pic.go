package main

import (
    "image/color"
    "image"
    "golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{125, 65, 255, 255}
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 100, 100)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
