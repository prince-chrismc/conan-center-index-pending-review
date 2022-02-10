package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
)

func MakeGif(images []image.Image, delay int) gif.GIF {
	// Alloc slice with 0 elems but capacity of all images
	frames := make([]*image.Paletted, 0, len(images)+1)
	delays := make([]int, 0, len(images)+1)
	disposals := make([]byte, 0, len(images)+1)

	for _, png := range images {
		// Images are in reverse order so we need to prepend them
		frames = append([]*image.Paletted{renderToPalette(png)}, frames...)
		delays = append(delays, delay)
		disposals = append(disposals, gif.DisposalBackground)
	}

	// Add out background to revert to when changing frames
	frames = append([]*image.Paletted{makeBlank()}, frames...)
	delays = append([]int{0}, delays...)
	disposals = append([]byte{gif.DisposalNone}, disposals...)

	return gif.GIF{
		Image:           frames,
		Delay:           delays,
		LoopCount:       10,
		Disposal:        disposals,
		BackgroundIndex: 0,
	}
}

func makeBlank() *image.Paletted {
	var palette color.Palette = color.Palette{
		image.Transparent,
	}
	img := image.NewPaletted(image.Rect(0, 0, 4025, 2048), palette)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0}}, img.Bounds().Min, draw.Src)
	return img
}

func renderToPalette(img image.Image) *image.Paletted {
	var palette color.Palette = color.Palette{
		image.Transparent,
		// color.RGBA{255, 255, 255, 255},
		color.RGBA{88, 166, 255, 255},
		color.RGBA{63, 185, 80, 255},
		color.RGBA{248, 81, 73, 255},
		color.RGBA{163, 113, 247, 255},
		color.RGBA{134, 94, 201, 255},
	}
	// Some days the images are off by a column so we are just hard coding the fix for now
	// TODO(prince-chrismc) Make this more generic
	sp := img.Bounds().Min
	width := img.Bounds().Dx()
	sp.X = width - 4025

	paletted := image.NewPaletted(image.Rect(0, 0, 4025, 2048), palette)
	draw.Draw(paletted, image.Rect(0, 0, 4025, 2048), img, sp, draw.Src)
	return paletted
}
