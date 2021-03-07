package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/format"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"golang.org/x/oauth2"
)

const (
	BUMP = "Bump version"
)

// Run the analysis
func Run(token string, dryRun bool) error {

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
	fmt.Printf("Limit: %d \nRemaining %d \n", rateLimit.Limit, rateLimit.Remaining)

	repo, _, err := client.Repository.Get(context, "conan-io", "conan-center-index")
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n-----\n", repo)

	var stats stats.Stats
	var retval []*pending_review.PullRequestStatus
	opt := &github.PullRequestListOptions{
		Sort:      "created",
		Direction: "asc",
		ListOptions: github.ListOptions{
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

		// Handle Pagination: https://github.com/google/go-github#pagination
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	bytes, err := json.MarshalIndent(retval, "", "   ")
	if err != nil {
		fmt.Printf("Problem formating result to JSON %v\n", err)
		os.Exit(1)
	}

	// https://github.com/prince-chrismc/conan-center-index-pending-review/issues/5#issuecomment-754112342
	isDifferent, err := validateContentIsDifferent(context, client, string(bytes))
	if err != nil {
		fmt.Printf("Problem getting original issue content %v\n", err)
		os.Exit(1)
	}

	if !isDifferent {
		fmt.Println("the obtained content is identical to the new result.")
		return nil // The published results are the same, no need to update the table.
	}

	if dryRun {
		fmt.Println(string(bytes))
		return nil
	}

	var ready string
	if stats.Merge > 0 {
		ready = `

### :heavy_check_mark: Ready to Merge 

Currently **` + fmt.Sprint(stats.Merge) + `** pull request(s) is/are waiting to be merged :tada:

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + format.ReviewsToMarkdownRows(retval, true)
	}

	_, _, err = client.Issues.Edit(context, "prince-chrismc", "conan-center-index-pending-review", 1, &github.IssueRequest{
		Body: github.String(`## :sparkles: Pull Requests Pending Review Summary!

### :ballot_box_with_check: Selection Criteria:

- There has been at least one approval (at any point)
- No reviews and commited to in the last 24hrs
- No labels with exception to "bump version" and "docs"

#### Legend

:new: - adding a recipe which does not yet exist
:arrow_up: - a version bump (matches the label)
:memo: - modification to an existing recipe
:green_book: - a documentation change (matches the label)
:warning: - the merge commit status does **not** indicate success (only displayed when ready to merge)

### :nerd_face: Please Review! 

There are **` + fmt.Sprint(stats.Review) + `** pull requests currently under way :eyes:

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + format.ReviewsToMarkdownRows(retval, false) + ready + `

#### :bar_chart: Statistics

> :warning: These are just rough metrics counthing the labels and may not reflect the acutal state of pull requests

- Commit: ` + os.Getenv("GITHUB_SHA") + `
- PRs
   - Open: ` + fmt.Sprint(stats.Open) + `
   - Draft: ` + fmt.Sprint(stats.Draft) + `
   - Age: ` + duration.Duration(stats.Age) + `
- Labels
   - Stale: ` + fmt.Sprint(stats.Stale) + `
   - Failed: ` + fmt.Sprint(stats.Failed) + `
   - Blocked: ` + fmt.Sprint(stats.Blocked) + `
` +
			"\n\n<details><summary>Raw JSON data</summary>\n\n```json\n" + string(bytes) + "\n```\n\n</details>"),
	})
	if err != nil {
		fmt.Printf("Problem editing issue %v\n", err)
		os.Exit(1)
	}

	return nil
}

func validateContentIsDifferent(context context.Context, client *pending_review.Client, expected string) (bool, error) {
	issue, _, err := client.Issues.Get(context, "prince-chrismc", "conan-center-index-pending-review", 1)
	if err != nil {
		return false, err
	}
	content := issue.GetBody()

	rawStart := strings.Index(content, "```json\n")
	rawEnd := strings.Index(content, "\n```\n")

	if rawStart == -1 || rawEnd == -1 {
		// Second chance... Editing the issues manually on windows will add CR to the entire content
		rawStart = strings.Index(content, "```json\r\n")
		rawEnd = strings.Index(content, "\r\n```\r\n")

		if rawStart == -1 || rawEnd == -1 {
			return false, errors.New("content did not contain the expected raw JSON section")
		}
	}

	obtained := content[rawStart+len("```json\n") : rawEnd]

	return obtained != expected, nil
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*pending_review.PullRequest) ([]*pending_review.PullRequestStatus, stats.Stats) {
	var stats stats.Stats
	var out []*pending_review.PullRequestStatus
	for _, pr := range prs {
		stats.Age = time.Duration(int64(float64((time.Now().Sub(pr.GetCreatedAt()) + time.Duration(stats.Age.Nanoseconds()*int64(stats.Open))).Nanoseconds()) / float64(stats.Open+1)))
		stats.Open++

		if pr.GetDraft() {
			stats.Draft++
			continue // Let's skip these
		}

		isBump := false
		isDoc := false
		len := len(pr.Labels)
		if len > 0 {
			for _, label := range pr.Labels {
				switch label.GetName() {
				case BUMP:
					isBump = true
				case "Docs":
					isDoc = true
				case "stale":
					stats.Stale++
				case "Failed", "Unexpected Error":
					stats.Failed++
				case "infrastructure", "blocked":
					stats.Blocked++
				}
			}

			if !isBump && !isDoc {
				continue // We know if there are certain labels then there's probably something wrong!
			}
		}

		if len > 1 && !isDoc { // We always want to review documentation changes
			continue // We know if there are certain labels then it's probably something worth skipping!
		}

		review, _, err := client.PullRequest.GatherRelevantReviews(context, "conan-io", "conan-center-index", pr)
		if errors.Is(err, pending_review.ErrNoReviews) || errors.Is(err, pending_review.ErrInvalidChange) {
			continue
		} else if err != nil {
			fmt.Printf("Problem getting list of reviews %v\n", err)
			os.Exit(1)
		}

		if isBump {
			review.Change = pending_review.BUMP // FIXME: It would be nice for this logic to be internal
		}

		if review.IsMergeable {
			stats.Merge++
		} else {
			stats.Review++
		}

		fmt.Printf("%+v\n", review)
		out = append(out, review)
	}
	return out, stats
}
