package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/charts"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"github.com/wcharczuk/go-chart/v2"
)

// TimeInReview analysis of merged pull requests
func TimeInReview(token string, dryRun bool, owner string, repo string) error {
	context := context.Background()
	client, err := internal.MakeClient(context, token, pending_review.TargetRepository{Owner: owner, Name: repo})
	if err != nil {
		fmt.Printf("Problem getting rate limit information %v\n", err)
		os.Exit(1)
	}

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	tir := make(stats.DurationAtTime) // Time in review
	mpd := make(stats.CountAtTime)    // Merged Per Day

	opt := &github.PullRequestListOptions{
		Sort:  "created",
		State: "closed",
		ListOptions: github.ListOptions{
			// Through browsing GitHub this is about where the meaningful data starts
			Page:    20,
			PerPage: 100,
		},
	}
	for {
		pulls, resp, err := client.PullRequests.List(context, "conan-io", "conan-center-index", opt)
		if err != nil {
			fmt.Printf("Problem getting pull request list %v\n", err)
			os.Exit(1)
		}

		for _, pull := range pulls {
			// The 'community reviewers' was fully emplace on Sept 28th 2020
			// see https://github.com/conan-io/conan-center-index/issues/2857#issuecomment-696221003
			if pull.GetClosedAt().Before(time.Date(2020, time.September, 28, 0, 0, 0, 0, time.UTC)) {
				continue
			}

			// These typically take little to no time and are sometimes forces through
			// https://github.com/conan-io/conan-center-index/pulls?q=is%3Apr+is%3Amerged+label%3ADocs
			if len(pull.Labels) > 0 && pull.Labels[0].GetName() == "Docs" {
				continue
			}

			merged := pull.GetMergedAt() != time.Time{} // `merged` is not returned when paging through the API - so calculate it
			if merged {
				fmt.Printf("#%4d was closed at %s and merged at %s\n", pull.GetNumber(), pull.GetClosedAt().String(), pull.GetMergedAt().String())
				tir[pull.GetMergedAt()] = pull.GetMergedAt().Sub(pull.GetCreatedAt())
				mpd.Count(pull.GetMergedAt().Truncate(time.Hour * 24))
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	lineGraph := charts.MakeLineChart(tir, mpd)

	if dryRun {
		f, _ := os.Create("tir.png")
		defer f.Close()
		err = lineGraph.Render(chart.PNG, f)
		if err != nil {
			fmt.Printf("Problem rendering %s %v\n", "tir.png", err)
			os.Exit(1)
		}

		return nil
	}

	_, err = internal.UpdateJSONFile(context, client, "time-in-review.json", tir)
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "time-in-review.json", err)
		os.Exit(1)
	}

	_, err = internal.UpdateJSONFile(context, client, "closed-per-day.json", mpd) // Legacy file name
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "closed-per-day.json", err) // Legacy file name
		os.Exit(1)
	}

	var b bytes.Buffer
	err = lineGraph.Render(chart.PNG, &b)
	if err != nil {
		fmt.Printf("Problem rendering %s %v\n", "time-in-review.png", err)
		os.Exit(1)
	}

	_, err = internal.UpdateDataFile(context, client, "time-in-review.png", b.Bytes())
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "time-in-review.png", err)
		os.Exit(1)
	}

	fmt.Println("::endgroup")

	return nil
}
