package charts

import (
	"fmt"
	"sort"
	"time"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func inReviewKeys(d stats.DurationAtTime) []time.Time {
	v := make([]time.Time, len(d))
	idx := 0
	for time := range d {
		v[idx] = time
		idx++
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i].Before(v[j])
	})

	return v
}

func inReviewDurationValues(d stats.DurationAtTime, sorted []time.Time) []float64 {
	v := make([]float64, len(d))
	idx := 0
	for _, value := range sorted {
		v[idx] = d[value].Hours() / 24.0
		idx++
	}
	return v
}

func closedCountValues(d stats.CountAtTime, sorted []time.Time) []float64 {
	v := make([]float64, len(d))
	idx := 0
	for _, value := range sorted {
		v[idx] = float64(d[value])
		idx++
	}
	return v
}

func MakeLineChart(tir stats.DurationAtTime, cpd stats.CountAtTime) chart.Chart {
	sortedData := inReviewKeys(tir)
	mainSeries := chart.TimeSeries{
		Name: "Time in review",
		Style: chart.Style{
			StrokeColor: chart.ColorBlue.WithAlpha(125),
			FillColor:   chart.ColorBlue.WithAlpha(50),
		},
		XValues: sortedData,
		YValues: inReviewDurationValues(tir, sortedData),
	}

	smaSeries := &chart.SMASeries{
		Name: "Moving average",
		Style: chart.Style{
			StrokeColor: drawing.ColorRed,
		},
		InnerSeries: mainSeries,
		Period:      75,
	}

	sortedTime := cpd.Keys()
	secondSeries := chart.TimeSeries{
		Name: "Closed per day",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("E5934C").WithAlpha(150),
		},
		YAxis:   chart.YAxisSecondary,
		XValues: sortedTime,
		YValues: closedCountValues(cpd, sortedTime),
	}

	graph := chart.Chart{
		Background: chart.Style{Padding: chart.Box{Top: 25, Left: 20}},
		XAxis: chart.XAxis{
			Name: "Closed At",
		},
		YAxis: chart.YAxis{
			Name: "Days",
			Ticks: []chart.Tick{
				{Value: 0.0, Label: "0"},
				{Value: 7.0, Label: "7"},
				{Value: 15.0, Label: "15"},
				{Value: 30.0, Label: "30"},
				{Value: 45.0, Label: "45"},
				{Value: 60.0, Label: "60"},
				{Value: 75.0, Label: "75"},
				{Value: 90.0, Label: "90"},
			},
			GridMajorStyle: chart.Style{
				StrokeColor:     chart.ColorAlternateGray,
				StrokeWidth:     1,
				StrokeDashArray: []float64{10.0, 25.0},
			},
			GridLines: []chart.GridLine{
				{Value: 7},
				{Value: 15},
				{Value: 30},
				{Value: 60},
			},
		},
		YAxisSecondary: chart.YAxis{
			Name: "Pull Requests",
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%.1f", vf)
				}
				return ""
			},
		},
		Series: []chart.Series{
			mainSeries,
			smaSeries,
			chart.LastValueAnnotationSeries(smaSeries),
			secondSeries,
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendThin(&graph),
	}

	return graph
}
