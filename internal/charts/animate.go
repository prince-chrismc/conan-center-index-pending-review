package charts

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
)

func MakeGif(images []image.Image, delay int) gif.GIF {
	// Alloc slice with 0 elems but capacity of all images
	frames := make([]*image.Paletted, 0, len(images))
	delays := make([]int, 0, len(images))

	for _, png := range images {
		frames = append([]*image.Paletted{renderToPalette(png)}, frames...)
		delays = append(delays, delay)
	}

	return gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: 10,
	}
}

func renderToPalette(img image.Image) *image.Paletted {
	var palette color.Palette = color.Palette{
		image.Transparent,
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

	paletted := image.NewPaletted(img.Bounds(), palette)
	draw.Draw(paletted, img.Bounds(), img, sp, draw.Over)
	return paletted
}
