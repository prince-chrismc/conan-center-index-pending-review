package main

import (
	"fmt"
	"os"
	"sort"
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

func toDays(raw []dataPoint) []float64 {
	var seconds []float64
	for _, d := range raw {
		seconds = append(seconds, d.Duration.Hours()/24.0)
	}
	return seconds
}

func arrayOfTime(cpd closedPerDay) []time.Time {
	v := make([]time.Time, len(cpd), len(cpd))
	idx := 0
	for time := range cpd {
		v[idx] = time
		idx++
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i].Before(v[j])
	})

	return v
}

func arrayOfCounts(cpd closedPerDay, sorted []time.Time) []float64 {
	v := make([]float64, len(cpd), len(cpd))
	idx := 0
	for _, value := range sorted {
		v[idx] = float64(cpd[value])
		idx++
	}
	return v
}

func makeChart(data []dataPoint, cpd closedPerDay) {

	mainSeries := chart.TimeSeries{
		Name: "Time in review",
		Style: chart.Style{
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: getCreatedAt(data),
		YValues: toDays(data),
	}

	smaSeries := &chart.SMASeries{
		Name: "Moving average of time",
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
			StrokeWidth:     2,
		},
		InnerSeries: mainSeries,
	}

	sortedTime := arrayOfTime(cpd)
	secondSeries := chart.TimeSeries{
		Name: "Closed per day",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("E5934C"),
		},
		YAxis:   chart.YAxisSecondary,
		XValues: sortedTime,
		YValues: arrayOfCounts(cpd, sortedTime),
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Left: 175,
			},
		},
		Title: "Summary of Time in Review",
		TitleStyle: chart.Style{
			Padding: chart.Box{
				Left: 175,
			},
		},
		XAxis: chart.XAxis{
			Name: "Created At",
		},
		YAxis: chart.YAxis{
			Name: "Days until Merged",
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					vi := int64(vf)
					return fmt.Sprintf("%d", vi)
				}
				return ""
			},
		},
		YAxisSecondary: chart.YAxis{
			Name: "Number of PRs Merged",
			NameStyle: chart.Style{
				Padding: chart.Box{
					Left: -20,
				},
			},
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
			secondSeries,
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
