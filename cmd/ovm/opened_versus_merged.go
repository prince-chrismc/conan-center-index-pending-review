package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/charts"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"github.com/urfave/cli/v2"
	"github.com/wcharczuk/go-chart/v2"
)

const interval = duration.WEEK * 52
const delay = 75

// OpenVersusMerged generates a graph depicting the last 1 year of pull requests highlighting where are open, close, and merged
func OpenVersusMerged(dryRun bool, c *cli.Context) error {
	context := context.Background()
	client, err := internal.MakeClient(context, c)
	if err != nil {
		return fmt.Errorf("problem making client %w", err)
	}

	opw := make(stats.CountAtTime)  // Opened Per Week
	cxw := make(stats.CountAtTime)  // Closed (based on creation date) Per Week
	mxw := make(stats.CountAtTime)  // Merged (based on creation date) Per Week
	m7xw := make(stats.CountAtTime) // Merged within 7 days (based on creation date) Per Week

	defer fmt.Println("::endgroup") // Always print when we return

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	err = countClosedPullRequests(context, client, opw, cxw, mxw, m7xw)
	if err != nil {
		return fmt.Errorf("problem counting closed pull requests %w", err)
	}

	err = countOpenedPullRequests(context, client, opw)
	if err != nil {
		return fmt.Errorf("problem counting open pull requests %w", err)
	}

	images, err := GetOvmPngFromThisWeek(context, client)
	if err != nil || len(images) == 0 { // We know there should always be commits
		return fmt.Errorf("problem getting %s history %w", "ovm.png", err)
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
		return fmt.Errorf("problem rendering %s %w", "open-versus-merged.png", err)
	}

	_, err = internal.UpdateDataFile(context, client, "open-versus-merged.png", b1.Bytes())
	if err != nil {
		return fmt.Errorf("problem updating %s %w", "open-versus-merged.png", err)
	}

	b2 := bytes.NewBuffer(b1.Bytes())
	img, err := png.Decode(b2)
	if err != nil {
		return fmt.Errorf("problem decoding %s %w", "ovm.png", err)
	}

	images = append([]image.Image{img}, images...)
	jif := MakeGif(images, delay)

	var b3 bytes.Buffer
	err = gif.EncodeAll(&b3, &jif)
	if err != nil {
		return fmt.Errorf("problem encoding %s %w", "ovm.gif", err)
	}

	_, err = internal.UpdateDataFile(context, client, "open-versus-merged.gif", b3.Bytes())
	if err != nil {
		return fmt.Errorf("problem updating %s %w", "open-versus-merged.gif", err)
	}

	return nil
}

func prCreationDay(pull *github.PullRequest) time.Time {
	return pull.GetCreatedAt().Truncate(duration.WEEK)
}

func countClosedPullRequests(context context.Context, client *pending_review.Client, opw stats.CountAtTime, cxw stats.CountAtTime, mxw stats.CountAtTime, m7xw stats.CountAtTime) error {
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
			return fmt.Errorf("problem getting pull request list %w", err)
		}

		for _, pull := range pulls {
			createdOn := prCreationDay(pull)
			if time.Since(createdOn) > interval {
				return nil
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
			return nil
		}
		opt.Page = resp.NextPage
	}
}

func countOpenedPullRequests(context context.Context, client *pending_review.Client, opw stats.CountAtTime) error {
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
			return fmt.Errorf("problem getting pull request list %w", err)

		}

		for _, pull := range pulls {
			createdOn := prCreationDay(pull)
			if time.Since(createdOn) > interval {
				return nil
			}

			opw.Count(createdOn)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return nil
}
