package charts

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func MakeStackedChart(opd stats.CountAtTime, cxd stats.CountAtTime, mxd stats.CountAtTime) chart.StackedBarChart {
	bars := []chart.StackedBar{}
	for _, t := range opd.Keys() {
		bars = append(bars, chart.StackedBar{
			Name: t.Format(chart.DefaultDateFormat),
			// Width: 25,
			Values: []chart.Value{
				{Value: float64(opd[t] - cxd[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("3fb950")}},
				{Value: float64(cxd[t] - mxd[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("f85149")}},
				{Value: float64(mxd[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("a371f7")}},
			},
		})
	}

	return chart.StackedBarChart{
		// Title:  "Open versus merged pull requests",
		Canvas: chart.Style{ /*Padding: chart.Box{Top: 40, Left: 40, Right: 40, Bottom: 40},*/ FillColor: chart.ColorWhite.WithAlpha(0)},
		// Background: chart.Style{Padding: chart.Box{Top: 100, Left: 100, Right: 100, Bottom: 100}, FillColor: chart.ColorWhite.WithAlpha(0)},
		Bars:       bars,
		BarSpacing: 25,
		Height:     2048,
		Width:      len(bars)*75 + 40,
	}
}
