package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func inReviewKeys(d timeInReview) []time.Time {
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

func inReviewDurationValues(d timeInReview, sorted []time.Time) []float64 {
	v := make([]float64, len(d))
	idx := 0
	for _, value := range sorted {
		v[idx] = d[value].Hours() / 24.0
		idx++
	}
	return v
}

func closedKeys(d closedPerDay) []time.Time {
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

func closedCountValues(d closedPerDay, sorted []time.Time) []float64 {
	v := make([]float64, len(d))
	idx := 0
	for _, value := range sorted {
		v[idx] = float64(d[value])
		idx++
	}
	return v
}

func makeChart(tir timeInReview, cpd closedPerDay) chart.Chart {
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
		Name: "Moving average of time",
		Style: chart.Style{
			StrokeColor: drawing.ColorRed,
		},
		InnerSeries: mainSeries,
		Period:      50,
	}

	sortedTime := closedKeys(cpd)
	secondSeries := chart.TimeSeries{
		Name: "Closed per day",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("E5934C").WithAlpha(150),
		},
		YAxis:   chart.YAxisSecondary,
		XValues: sortedTime,
		YValues: closedCountValues(cpd, sortedTime),
	}

	padding := chart.Style{Padding: chart.Box{Left: 175}}
	graph := chart.Chart{
		Background: padding,
		Title:      "Time Spent in Review",
		TitleStyle: padding,
		XAxis: chart.XAxis{
			Name: "Closed At",
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
			GridMajorStyle: chart.Style{
				StrokeColor:     chart.ColorAlternateGray,
				StrokeWidth:     1,
				StrokeDashArray: []float64{10.0, 25.0},
			},
			GridLines: []chart.GridLine{
				{Value: 15},
				{Value: 30},
				{Value: 60},
			},
		},
		YAxisSecondary: chart.YAxis{
			Name: "PRs Merged",
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
		chart.LegendLeft(&graph),
	}

	return graph
}
