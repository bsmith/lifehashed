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
			// 1. - 2.^-Floor[5.*x]
			z := (0.5 + float64(x)) / float64(width)
			linear := 1. - math.Pow(2., -math.Floor(5.*z))
			linear32 := float32(linear)
			srgb := srgb.ColorFromLinear(linear32, linear32, linear32)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 2 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 1. - 2.^-Floor[5.*x]
			z := (0.5 + float64(x)) / float64(width)
			linear := 1. - math.Pow(2., -5.*z)
			linear32 := float32(linear)
			srgb := srgb.ColorFromLinear(linear32, linear32, linear32)
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	for y := 3 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 1. - 2.^-Floor[5.*x]
			z := (0.5 + float64(x)) / float64(width)
			linear := 1. - math.Pow(2., -5.*z)
			srgb := color.RGBA{uint8(255 * linear), uint8(255 * linear), uint8(255 * linear), 255}
			image.Set(x, y, srgb)
		}
	}

	for y := 4 * bandSize; y < height; y++ {
		for x := 0; x < width; x++ {
			// 1. - 2.^-Floor[5.*x]
			z := (0.5 + float64(x)) / float64(width)
			linear := 1. - math.Pow(2., -5.*z)
			// linear32 := float32(linear)
			// srgb := srgb.ColorFromXYZ(ciexyz.Color{X: linear32, Y: linear32, Z: linear32})
			srgb := srgb.ColorFromXYZ(ciexyz.ColorFromV(ciexyz.D65.ToV().MulS(linear)))
			image.Set(x, y, srgb.ToRGBA(1.0))
		}
	}

	writer, err := os.Create("life01.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(writer, image)
}
