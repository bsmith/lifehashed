package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/srgb"
)

func main() {
	width := 64
	height := 64
	bandSize := 8
	stops := 8.
	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// level := uint8(math.Pow((0.5+float64(x))/float64(width), 1/2.2) * 255)
			linear := (0.5 + float32(x)) / float32(width)
			srgb := srgb.ColorFromLinear(linear, linear, linear)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 1 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -math.Floor(stops*(1-z)))
			linear32 := float32(linear)
			srgb := srgb.ColorFromLinear(linear32, linear32, linear32)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 2 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			linear32 := float32(linear)
			srgb := srgb.ColorFromLinear(linear32, linear32, linear32)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 3 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			srgb := color.RGBA{uint8(255 * linear), uint8(255 * linear), uint8(255 * linear), 255}
			image.Set(x, y, srgb)
		}
	}

	for y := 4 * bandSize; y < height; y++ {
		colourV := ciexyz.D65.ToV()
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			srgb := srgb.ColorFromXYZ(ciexyz.ColorFromV(colourV.MulS(linear)))
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 5 * bandSize; y < height; y++ {
		colourV := ciexyz.ColorFromXYY(srgb.PrimaryRed).ToV()
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			srgb := srgb.ColorFromXYZ(ciexyz.ColorFromV(colourV.MulS(linear)))
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 6 * bandSize; y < height; y++ {
		colourV := ciexyz.ColorFromXYY(srgb.PrimaryGreen).ToV()
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			srgb := srgb.ColorFromXYZ(ciexyz.ColorFromV(colourV.MulS(linear)))
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 7 * bandSize; y < height; y++ {
		colourV := ciexyz.ColorFromXYY(srgb.PrimaryBlue).ToV()
		for x := 0; x < width; x++ {
			// 2.^-Floor[5.*(1 - x)]
			z := (0.5 + float64(x)) / float64(width)
			linear := math.Pow(2., -stops*(1-z))
			srgb := srgb.ColorFromXYZ(ciexyz.ColorFromV(colourV.MulS(linear)))
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	writer, err := os.Create("demo01.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(writer, image)
}
