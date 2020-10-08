package main

import(
	"github.com/JiovaniLedesma/image_processing/image_processing"
)

func main() {
	img, err := image_processing.OpenImage("/home/jiovani/Pictures/Wallpapers/abstract.png")
	if err != nil {
		return
	}
}
