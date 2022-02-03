package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

// Save to Disk write the bar graph to PNG and the weeks results to GIF
func SaveToDisk(barGraph chart.StackedBarChart, images []image.Image) error {
	var b bytes.Buffer
	err := barGraph.Render(chart.PNG, &b)
	if err != nil {
		fmt.Printf("Problem rendering %s %v\n", "ovm.png", err)
		return err
	}

	f, _ := os.Create("ovm.png")
	defer f.Close()
	f.Write(b.Bytes())

	img, err := png.Decode(&b)
	if err != nil {
		fmt.Printf("Problem decoding %s %v\n", "ovm.png", err)
		return err
	}

	// Alloc slice with 0 elems but capacity of all previous images + new latest image
	frames := make([]*image.Paletted, 0, len(images)+1)
	delays := make([]int, 0, len(images)+1)

	// TODO(prince-chrismc) The last one is placed weirdly...
	for _, png := range images[:len(images)-1] {
		frames = append([]*image.Paletted{renderToPalette(png)}, frames...)
		delays = append(delays, delay)
	}

	lastFrame := renderToPalette(img)
	frames = append(frames, lastFrame)
	delays = append(delays, delay)

	jif := gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: 10,
	}

	g, _ := os.Create("ovm.gif")
	defer g.Close()

	err = gif.EncodeAll(g, &jif)
	if err != nil {
		fmt.Printf("Problem encoding %s %v\n", "ovm.gif", err)
		return err
	}

	return nil
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
	paletted := image.NewPaletted(img.Bounds(), palette)
	draw.Draw(paletted, img.Bounds(), img, img.Bounds().Min, draw.Over)
	return paletted
}
