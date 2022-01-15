package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"time"

	"github.com/google/go-github/v42/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/charts"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"github.com/wcharczuk/go-chart/v2"
	"golang.org/x/oauth2"
)

const interval = duration.WEEK * 10

// OpenVersusMerged generates a graph depicting the last 1 year of pull requests highlighting where are open, close, and merged
func OpenVersusMerged(token string, dryRun bool) error {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	context := context.Background()
	client := pending_review.NewClient(oauth2.NewClient(context, tokenService))

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		fmt.Printf("Problem getting rate limit information %v\n", err)
		os.Exit(1)
	}

	// We have not exceeded the limit so we can continue
	fmt.Printf("Limit: %d \nRemaining: %d \n", rateLimit.Limit, rateLimit.Remaining)

	opw := make(stats.CountAtTime)  // Opend Per Week
	cxw := make(stats.CountAtTime)  // Closed (based on creation date) Per Week
	mxw := make(stats.CountAtTime)  // Merged (based on creation date) Per Week
	m7xw := make(stats.CountAtTime) // Merged within 7 days (based on creation date) Per Week

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	countClosedPullRequests(context, client, opw, cxw, mxw, m7xw)
	countOpenedPullRequests(context, client, opw)

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	barGraph := charts.MakeStackedChart(opw, cxw, mxw, m7xw)

	if dryRun {
		f, _ := os.Create("ovm.png")
		defer f.Close()
		err = barGraph.Render(chart.PNG, f)
		if err != nil {
			fmt.Printf("Problem rendering %s %v\n", "ovm.png", err)
			os.Exit(1)
		}

		f1, _ := os.Open("ovm.png")
		defer f1.Close()
		img, err := png.Decode(f1)
		if err != nil {
			fmt.Printf("Problem decoding %s %v\n", "ovm.png", err)
			os.Exit(1)
		}

		var palette color.Palette = color.Palette{
			image.Transparent,
			color.RGBA{88, 166, 255, 255},
			color.RGBA{63, 185, 80, 255},
			color.RGBA{248, 81, 73, 255},
			color.RGBA{163, 113, 247, 255},
			color.RGBA{134, 94, 201, 255},
		}
		empty := image.NewPaletted(img.Bounds(), palette)
		draw.Draw(empty, img.Bounds(), img, img.Bounds().Min, draw.Over)

		// gif.Encode(palette, img, &gif.Options{
		// 	Quantizer: draw.Quantizer,
		// })

		jif := gif.GIF{
			Image: []*image.Paletted{empty},
			Delay: []int{0},
		}

		g, _ := os.Create("ovm.gif")
		defer g.Close()

		err = gif.EncodeAll(g, &jif)
		if err != nil {
			fmt.Printf("Problem encoding %s %v\n", "ovm.gif", err)
			os.Exit(1)
		}

		return nil
	}

	var b bytes.Buffer
	err = barGraph.Render(chart.PNG, &b)
	if err != nil {
		fmt.Printf("Problem rendering %s %v\n", "open-versus-merged.png", err)
		os.Exit(1)
	}

	_, err = internal.UpdateDataFile(context, client, "open-versus-merged.png", b.Bytes())
	if err != nil {
		fmt.Printf("Problem updating %s %v\n", "open-versus-merged.png", err)
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
