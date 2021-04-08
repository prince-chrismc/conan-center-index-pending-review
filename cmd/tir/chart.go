package main

import (
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func getCreatedAt(data []dataPoint) []time.Time {
	var dates []time.Time
	for i, ts := range data {
		dates = append(dates, ts.Created)
		if !dates[i].Equal(ts.Created) {
			panic("values do not match!")
		}
	}
	return dates
}

func toHours(raw []dataPoint) []float64 {
	var seconds []float64
	for _, d := range raw {
		seconds = append(seconds, d.Duration.Hours())
	}
	return seconds
}

func makeChart(data []dataPoint) {

	mainSeries := chart.TimeSeries{
		Name: "Time in review",
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: getCreatedAt(data),
		YValues: toHours(data),
	}

	// secondSeries := chart.TimeSeries{
	// 	Name: "SPY",
	// 	Style: chart.Style{
	// 		StrokeColor: drawing.ColorFromHex("efefef"),
	// 		FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
	// 	},
	// 	XValues: shortXvalues(),
	// 	YValues: shortYvalues(),
	// }

	// note we create a SimpleMovingAverage series by assignin the inner series.
	// we need to use a reference because `.Render()` needs to modify state within the series.
	smaSeries := &chart.SMASeries{
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	} // we can optionally set the `WindowSize` property which alters how the moving average is calculated.

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Created At",
		},
		YAxis: chart.YAxis{
			Name: "Hours until Merged",
		},
		Series: []chart.Series{
			mainSeries,
			smaSeries,
			// secondSeries,
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
