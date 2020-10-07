package main

import "example_image_processing.go"

func main() {
	img, err := image_processing.openImage("/home/jiovani/Pictures/Wallpapers/abstract.png")
	if err != nil {
		return
	}
}
