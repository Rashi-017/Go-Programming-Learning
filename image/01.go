package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	width := 100
	height := 100

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill it with red color
	red := color.RGBA{255, 255, 0, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, red)
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, red)
		}
	}

	// Save to file
	f, err := os.Create("red.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img) // Save as PNG
}
