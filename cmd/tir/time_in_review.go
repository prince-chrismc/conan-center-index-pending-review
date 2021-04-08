package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"golang.org/x/oauth2"
)

type dataPoint struct {
	Created  time.Time
	Duration time.Duration
}

type closedPerDay map[time.Time]int

// TimeInReview analysis of merged pull requests
func TimeInReview(token string, dryRun bool) error {
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

	var retval []dataPoint
	cpd := make(closedPerDay)
	opt := &github.PullRequestListOptions{
		Sort:  "created",
		State: "closed",
		ListOptions: github.ListOptions{
			Page:    22, // Through browsing GitHub this is about where the meaningful data starts
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
			// The 'community reviewers' was fully emplace on Sept 28th 2020, however it seems to have taken a little longer to see the effects
			// see https://github.com/conan-io/conan-center-index/issues/2857#issuecomment-696221003
			if pull.GetCreatedAt().Before(time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC)) {
				continue
			}

			merged := pull.GetMergedAt() != time.Time{} // merged is not returned when paging through the API - so calculated
			if merged {
				fmt.Printf("#%4d was created at %s and merged at %s\n", pull.GetNumber(), pull.GetCreatedAt().String(), pull.GetMergedAt().String())
				retval = append(retval, dataPoint{
					pull.GetCreatedAt(),
					pull.GetMergedAt().Sub(pull.GetCreatedAt()),
				})
				mergedOn := pull.GetMergedAt().Truncate(time.Hour * 24)
				currentCounter, found := cpd[mergedOn]
				if found {
					cpd[mergedOn] = currentCounter + 1
				} else {
					cpd[mergedOn] = 1
				}
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage

	}

	// sort.SliceStable(retval, func(i, j int) bool {
	// 	return retval[i].Created.Before(retval[j].Created)
	// })

	makeChart(retval, cpd)

	// bytes, err := json.MarshalIndent(retval, "", "   ")
	// if err != nil {
	// 	fmt.Printf("Problem formating result to JSON %v\n", err)
	// 	os.Exit(1)
	// }

	// commentBody := `### :see_no_evil: Raw date for time in review!

	// ` + "```json\n" + string(bytes) + "\n```"

	// if dryRun {
	// 	// fmt.Println(commentBody)
	// 	return nil
	// }

	return nil
}
