package main

import(
	"github.com/JiovaniLedesma/image_processing/img_processing"
)

func main() {
	img, err := img_processing.openImage("/home/jiovani/Pictures/Wallpapers/abstract.png")
	if err != nil {
		return
	}
}
