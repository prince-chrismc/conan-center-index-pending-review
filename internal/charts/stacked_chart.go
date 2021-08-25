package charts

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func MakeStackedChart(opd stats.CountAtTime, cxd stats.CountAtTime, mxd stats.CountAtTime, m7xw stats.CountAtTime) chart.StackedBarChart {
	most := opd.Values()[0]
	bars := []chart.StackedBar{}
	for _, t := range opd.Keys() {
		bars = append(bars, chart.StackedBar{
			Name: t.Format(chart.DefaultDateFormat),
			// Width: 25,
			Values: []chart.Value{
				{Value: float64(most - opd[t]), Style: chart.Style{FillColor: chart.ColorWhite.WithAlpha(0), StrokeColor: chart.ColorWhite.WithAlpha(0)}},
				{Value: float64(opd[t] - cxd[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("3fb950"), StrokeColor: chart.ColorWhite.WithAlpha(0)}},
				{Value: float64(cxd[t] - mxd[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("f85149"), StrokeColor: chart.ColorWhite.WithAlpha(0)}},
				{Value: float64(mxd[t] - m7xw[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("a371f7"), StrokeColor: chart.ColorWhite.WithAlpha(0)}},
				{Value: float64(m7xw[t]), Style: chart.Style{FillColor: drawing.ColorFromHex("865ec9"), StrokeColor: chart.ColorWhite.WithAlpha(0)}},
			},
		})
	}

	return chart.StackedBarChart{
		Title: "Open versus merged pull requests",
		TitleStyle: chart.Style{
			FontSize:          50,
			FontColor:         drawing.ColorFromHex("58a6ff"),
			TextVerticalAlign: chart.TextVerticalAlignTop,
		},
		Canvas:     chart.Style{FillColor: chart.ColorWhite.WithAlpha(0)},
		Background: chart.Style{Padding: chart.Box{Top: 125}},
		Bars:       bars,
		BarSpacing: 25,
		Height:     2048,
		Width:      len(bars)*75 + 40,
	}
}
