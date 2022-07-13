package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v42/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/format"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/validate"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// PendingReview analysis of open pull requests
func PendingReview(token string, dryRun bool) error {
	context := context.Background()
	client, err := internal.MakeClient(context, token)
	if err != nil {
		fmt.Printf("Problem getting rate limit information %v\n", err)
		os.Exit(1)
	}

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	var stats stats.Stats
	var retval []*pending_review.PullRequestSummary
	opt := &pending_review.PullRequestListOptions{
		Sort:      "created",
		Direction: "asc",
		ListOptions: pending_review.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		pulls, resp, err := client.PullRequests.List(context, "conan-io", "conan-center-index", opt)
		if err != nil {
			fmt.Printf("Problem getting pull request list %v\n", err)
			os.Exit(1)
		}

		out, s := gatherReviewStatus(context, client, pulls)
		retval = append(retval, out...)
		stats.Add(s)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	if !dryRun {
		isDifferent, err := internal.UpdateJSONFile(context, client, "pending-review.json", retval)
		if err != nil {
			fmt.Printf("Problem updating 'pending-review.json' %v\n", err)
			os.Exit(1)
		}

		// https://github.com/prince-chrismc/conan-center-index-pending-review/issues/5#issuecomment-754112342
		if !isDifferent {
			fmt.Println("🦺 The obtained content is identical to the new result.")
			return nil // The published results are the same, no need to update the table.
		}

		err = updateCountFile(context, client, "review-count.json", len(retval))
		if err != nil {
			fmt.Printf("Problem updating 'review-count.json' %v\n", err)
			os.Exit(1)
		}

		err = updateCountFile(context, client, "open-count.json", stats.Open)
		if err != nil {
			fmt.Printf("Problem updating 'open-count.json' %v\n", err)
			os.Exit(1)
		}
	}

	commentBody := `## :sparkles: Summary of Pull Requests Pending Review!

### :ballot_box_with_check: Selection Criteria:

- There has been at least one approval on the head commit
- The last commit occured after any reviews
- No labels with exception to "docs" and "GitHub config"

#### Legend

:new: - Adding a recipe which does not yet exist
:memo: - Modification to an existing recipe
:green_book: - Documentation change <sup>[1]</sup>
:gear: - GitHub configuration/workflow changes <sup>[1]</sup>
:stopwatch: or :warning: - The commit status does **not** indicate success <sup>[2]</sup>
:bell: - The last review was more than 12 days ago
:eyes: - It's been more than 3 days since the last commit and there are no reviews

<sup>[1]</sup>: _closely_ matches the label
<sup>[2]</sup>: depending whether the PR is under way or ready to merge` +
		format.UnderReview(retval) + format.ReadyToMerge(retval) + format.Statistics(stats) + `
		
[Raw JSON data](https://raw.githubusercontent.com/` + os.Getenv("GITHUB_REPOSITORY") + `/raw-data/pending-review.json)

## :bar_chart: Open Versus Merged

#### Legend

:green_square: - Open pull requests
:red_square: - Closed pull requests
:purple_square: - Merged pull requests <sup>[1]</sup>

![ovm](https://github.com/` + os.Getenv("GITHUB_REPOSITORY") + `/blob/raw-data/open-versus-merged.gif?raw=true)

<sup>[1]</sup>: the darker bottom section indicated merged within 7 days of being opened

## :hourglass: Time Spent in Review

![tir](https://github.com/` + os.Getenv("GITHUB_REPOSITORY") + `/blob/raw-data/time-in-review.png?raw=true)

Found this useful? Give it a :star: :pray:
`

	if dryRun {
		fmt.Println(commentBody)
		return nil
	}

	_, err = internal.UpdateFileAtRef(context, client, "index.md", "gh-pages", []byte(commentBody))
	if err != nil {
		fmt.Printf("Problem editing web view %v\n", err)
		os.Exit(1)
	}

	_, _, err = client.Issues.Edit(context, os.Getenv("GITHUB_REPOSITORY_OWNER"), "conan-center-index-pending-review", 1, &github.IssueRequest{
		Body: github.String(commentBody),
	})
	if err != nil {
		fmt.Printf("Problem editing issue %v\n", err)
		os.Exit(1)
	}

	fmt.Println("::endgroup")

	return nil
}

func updateCountFile(context context.Context, client *pending_review.Client, file string, count int) error {
	counts := stats.CountAtTime{}
	err := internal.GetJSONFile(context, client, file, &counts)
	if err != nil {
		return err
	}

	counts.AddNow(count)

	_, err = internal.UpdateJSONFile(context, client, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*pending_review.PullRequest) ([]*pending_review.PullRequestSummary, stats.Stats) {
	var stats stats.Stats
	var out []*pending_review.PullRequestSummary
	for _, pr := range prs {
		stats.Age.Append(time.Since(pr.GetCreatedAt()))
		stats.Open++

		if pr.GetDraft() {
			stats.Draft++
			continue // Let's skip these
		}

		valid := validate.OnlyAcceptableLabels(pr.Labels, &stats)
		if !valid {
			continue
		}

		review, _, err := client.PullRequest.GetReviewSummary(context, "conan-io", "conan-center-index", pr)
		if errors.Is(err, pending_review.ErrNoReviews) || errors.Is(err, pending_review.ErrInvalidChange) {
			continue
		} else if err != nil {
			fmt.Printf("Problem getting list of reviews %v\n", err)
			os.Exit(1)
		}

		if review.Summary.IsApproved() {
			stats.Merge++
		} else {
			stats.Review++
		}

		fmt.Printf("%+v\n", review)
		out = append(out, review)
	}
	return out, stats
}
