package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/charts"
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
	_, err = f.Write(b.Bytes())
	if err != nil {
		fmt.Printf("Problem writting %s %v\n", "ovm.png", err)
		return err
	}

	img, err := png.Decode(&b)
	if err != nil {
		fmt.Printf("Problem decoding %s %v\n", "ovm.png", err)
		return err
	}

	images = append([]image.Image{img}, images...)
	jif := charts.MakeGif(images, delay)

	g, _ := os.Create("ovm.gif")
	defer g.Close()

	err = gif.EncodeAll(g, &jif)
	if err != nil {
		fmt.Printf("Problem encoding %s %v\n", "ovm.gif", err)
		return err
	}

	return nil
}
