package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v34/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"golang.org/x/oauth2"
)

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

	opw := make(stats.CountAtTime) // Opend Per Week
	cxw := make(stats.CountAtTime) // Closed (based on creation date) Per Day
	mxw := make(stats.CountAtTime) // Merged (based on creation date) Per Day

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	fn0(tokenService, context, cxw, mxw)
	fn1(tokenService, context, opw)

	fmt.Println("::endgroup")

	return nil
}

func fn0(tokenService oauth2.TokenSource, context context.Context, opw stats.CountAtTime, cxw stats.CountAtTime, mxw stats.CountAtTime) {
	client := pending_review.NewClient(oauth2.NewClient(context, tokenService))

	opt := &github.PullRequestListOptions{
		Sort:  "created",
		State: "closed",
		ListOptions: github.ListOptions{
			Page:    30,
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
			if time.Since(pull.GetCreatedAt()) > time.Hour*24*60 {
				continue
			}

			opw.Count(pull.GetCreatedAt().Truncate(time.Hour * 24))
			cxw.Count(pull.GetCreatedAt().Truncate(time.Hour * 24))

			merged := pull.GetMergedAt() != time.Time{}
			if merged {
				mxw.Count(pull.GetCreatedAt().Truncate(time.Hour * 24))
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}

func fn1(tokenService oauth2.TokenSource, context context.Context, opw stats.CountAtTime) {
	client := pending_review.NewClient(oauth2.NewClient(context, tokenService))

	opt := &github.PullRequestListOptions{
		Sort:  "created",
		State: "opened",
		ListOptions: github.ListOptions{
			Page:    30,
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
			if time.Since(pull.GetCreatedAt()) > time.Hour*24*60 {
				continue
			}

			opw.Count(pull.GetCreatedAt().Truncate(time.Hour * 24))
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}
