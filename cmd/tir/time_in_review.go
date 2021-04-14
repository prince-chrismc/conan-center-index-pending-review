package main

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"golang.org/x/oauth2"
)

type dataPoint map[time.Time]time.Duration
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

	retval := make(dataPoint)
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
				retval[pull.GetMergedAt()] = pull.GetMergedAt().Sub(pull.GetCreatedAt())
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

	makeChart(retval, cpd)

	bytes, err := json.MarshalIndent(retval, "", "   ")
	if err != nil {
		fmt.Printf("Problem formating result to JSON %v\n", err)
		os.Exit(1)
	}

	fileContent, _, _, err := client.Repositories.GetContents(context, "prince-chrismc", "conan-center-index-pending-review", "time-in-review.json", &github.RepositoryContentGetOptions{
		Ref: "raw-data",
	})
	if err != nil {
		fmt.Printf("Problem getting current file %v\n", err)
		os.Exit(1)
	}

	newSha := fmt.Sprint(sha1.Sum(bytes))
	if newSha != fileContent.GetSHA() {
		opts := &github.RepositoryContentFileOptions{
			SHA:       fileContent.SHA, // Required to edit the file
			Message:   github.String("Time in review: New data - " + time.Now().Format(time.RFC3339)),
			Content:   bytes,
			Branch:    github.String("raw-data"),
			Committer: &github.CommitAuthor{Name: github.String("github-actions[bot]"), Email: github.String("github-actions[bot]@users.noreply.github.com")},
		}
		_, _, err = client.Repositories.UpdateFile(context, "prince-chrismc", "conan-center-index-pending-review", "time-in-review.json", opts)
		if err != nil {
			fmt.Printf("Problem creating file %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Content for 'time-in-review.json' was the same")
	}

	return nil
}
