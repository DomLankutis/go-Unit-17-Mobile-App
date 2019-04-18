package Utils

import (
	"golang.org/x/mobile/asset"
	"image"
	_ "image/png"
	"log"
)

func OpenFile(path string) asset.File{
	file, err := asset.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func OpenImage(imgPath string) image.Image {
	img := OpenFile(imgPath)
	defer img.Close()

	imgData, _, err := image.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	return imgData
}