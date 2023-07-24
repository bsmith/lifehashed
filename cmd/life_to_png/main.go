package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/mandykoh/prism/srgb"
)

func main() {
	width := 64
	height := 64
	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// level := uint8(math.Pow((0.5+float64(x))/float64(width), 1/2.2) * 255)
			linear := (0.5 + float32(x)) / float32(width)
			srgb := srgb.ColorFromLinear(linear, linear, linear)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := height / 2; y < height; y++ {
		for x := 0; x < width; x++ {
			// level := uint8(math.Pow((0.5+float64(x))/float64(width), 1/2.2) * 255)
			linear := 1 - float32(math.Pow(2, -5*(0.5+float64(x))/float64(width)))
			srgb := srgb.ColorFromLinear(linear, linear, linear)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	writer, err := os.Create("life01.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(writer, image)
}
