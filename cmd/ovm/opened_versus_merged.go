package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/charts"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"github.com/wcharczuk/go-chart/v2"
)

const interval = duration.WEEK * 52
const delay = 75

// OpenVersusMerged generates a graph depicting the last 1 year of pull requests highlighting where are open, close, and merged
func OpenVersusMerged(token string, dryRun bool, owner string, repo string) error {
	context := context.Background()
	client, err := internal.MakeClient(context, token, pending_review.WorkingRepository{Owner: owner, Name: repo})
	if err != nil {
		fmt.Printf("Problem getting rate limit information %v\n", err)
		os.Exit(1)
	}

	opw := make(stats.CountAtTime)  // Opened Per Week
	cxw := make(stats.CountAtTime)  // Closed (based on creation date) Per Week
	mxw := make(stats.CountAtTime)  // Merged (based on creation date) Per Week
	m7xw := make(stats.CountAtTime) // Merged within 7 days (based on creation date) Per Week

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	countClosedPullRequests(context, client, opw, cxw, mxw, m7xw)
	countOpenedPullRequests(context, client, opw)

	images, err := GetOvmPngFromThisWeek(context, client)
	if err != nil || len(images) == 0 { // We know there should always be commits
		fmt.Printf("Problem getting %s history %v\n", "ovm.png", err)
		os.Exit(1)
	}

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	barGraph := charts.MakeStackedChart(opw, cxw, mxw, m7xw)

	if dryRun {
		err = SaveToDisk(barGraph, images)

		fmt.Println("::endgroup")
		return err
	}

	var b1 bytes.Buffer
	err = barGraph.Render(chart.PNG, &b1)
	if err != nil {
		fmt.Printf("Problem rendering %s %v\n", "open-versus-merged.png", err)
		os.Exit(1)
	}

	_, err = internal.UpdateDataFile(context, client, "open-versus-merged.png", b1.Bytes())
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "open-versus-merged.png", err)
		os.Exit(1)
	}

	b2 := bytes.NewBuffer(b1.Bytes())
	img, err := png.Decode(b2)
	if err != nil {
		fmt.Printf("Problem decoding %s %v\n", "ovm.png", err)
		return err
	}

	images = append([]image.Image{img}, images...)
	jif := MakeGif(images, delay)

	var b3 bytes.Buffer
	err = gif.EncodeAll(&b3, &jif)
	if err != nil {
		fmt.Printf("Problem encoding %s %v\n", "ovm.gif", err)
		os.Exit(1)
	}

	_, err = internal.UpdateDataFile(context, client, "open-versus-merged.gif", b3.Bytes())
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "open-versus-merged.gif", err)
		os.Exit(1)
	}

	fmt.Println("::endgroup")

	return nil
}

func prCreationDay(pull *github.PullRequest) time.Time {
	return pull.GetCreatedAt().Truncate(duration.WEEK)
}

func countClosedPullRequests(context context.Context, client *pending_review.Client, opw stats.CountAtTime, cxw stats.CountAtTime, mxw stats.CountAtTime, m7xw stats.CountAtTime) {
	opt := &github.PullRequestListOptions{
		Sort:      "created",
		State:     "closed",
		Direction: "desc",
		ListOptions: github.ListOptions{
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
			createdOn := prCreationDay(pull)
			if time.Since(createdOn) > interval {
				return
			}

			opw.Count(createdOn)
			cxw.Count(createdOn)

			mergedOn := pull.GetMergedAt()
			merged := mergedOn != time.Time{}
			if merged {
				mxw.Count(createdOn)
				if mergedOn.Sub(pull.GetCreatedAt()) < duration.WEEK {
					m7xw.Count(createdOn)
				}
			}
		}

		if resp.NextPage == 0 {
			return
		}
		opt.Page = resp.NextPage
	}
}

func countOpenedPullRequests(context context.Context, client *pending_review.Client, opw stats.CountAtTime) {
	opt := &github.PullRequestListOptions{
		Sort:      "created",
		State:     "opened",
		Direction: "desc",
		ListOptions: github.ListOptions{
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
			createdOn := prCreationDay(pull)
			if time.Since(createdOn) > interval {
				return
			}

			opw.Count(createdOn)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}
