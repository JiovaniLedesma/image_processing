package main

import (
	"image/color"

	"github.com/JiovaniLedesma/image_processing/image_processing"
)

func main() {
	img, err := image_processing.OpenImage("/home/jiovani/Pictures/Wallpapers/abstract.png")
	if err != nil {
		return
	}

	size := img.Bounds().Size()
	// var pixels [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
	}
}
