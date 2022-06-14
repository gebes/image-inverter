package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

const folder = "./images/"

func main() {
	err := convertAll()
	if err != nil {
		fmt.Println("Could not convert:", err)
	}
}

func convertAll() error {
	files, err := ioutil.ReadDir(folder + "src/")
	if err != nil {
		return err
	}
	for _, file := range files {
		err = convert(folder+"src/"+file.Name(), folder+"out/"+file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func convert(from, to string) error {
	data, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	invertedImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			R, G, B, _ := img.At(x, y).RGBA()
			invertedPixel := color.RGBA{
				R: uint8(255 - R),
				G: uint8(255 - G),
				B: uint8(255 - B),
			}
			invertedImg.Set(x, y, invertedPixel)
		}
	}

	f, err := os.Create(to)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, invertedImg)
}
