package image_processing

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println("Imagen leída: ", fi.Name())
	//defer f.Close()sss
	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}
	if format != "jpeg" {
		fmt.Println("image format is not jpeg")
		return nil, errors.New("")
	}
	return img, nil
}
