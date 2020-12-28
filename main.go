package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v1/pending_review"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

const (
	// Labels
	BUMP_VERSION = "Bump Version"
	UNEXP_ERR    = "Unexpected Error"
)

func main() {
	context := context.Background()
	client := pending_review.NewClient(determineAndSetupCredentials(context))

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		fmt.Printf("Problem getting rate limit information %v\n", err)
		return
	}

	// We have not exceeded the limit so we can continue
	fmt.Printf("Limit: %d \nRemaining %d \n", rateLimit.Limit, rateLimit.Remaining)

	repo, _, err := client.Repository.Get(context, "conan-io", "conan-center-index")
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n-----\n", repo)

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

		out := gatherReviewStatus(context, client, pulls)
		retval = append(retval, out...)

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

	_, _, err = client.Issues.Edit(context, "prince-chrismc", "conan-center-index-pending-review", 1, &github.IssueRequest{
		Body: github.String(`## :sparkles: Pull Requests Pending Review Summary!

### :ballot_box_with_check: Selection Criteria:

- No reviews and pull requests were created in the last 24hrs
- There has been at least one approval

### :nerd_face: Please Review!

Number | Opened By | Title | Reviews | :stop_sign: Blockers | :heavy_check_mark: Approvers :star2:
:---: | --- | --- | :---: | --- | ---
` + formatPullRequestToMarkdownRows(retval) + "\n\n<details><summary>Raw JSON data</summary>\n\n```json\n" + string(bytes) + "\n```\n\n</details>"),
	})
	if err != nil {
		fmt.Printf("Problem editing issue %v\n", err)
		os.Exit(1)
	}
}

func formatPullRequestToMarkdownRows(prs []*pending_review.PullRequestStatus) string {
	var retval string
	for _, pr := range prs {
		column := []string{
			fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
			fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
			fmt.Sprintf("%.12q", pr.Title),
			fmt.Sprint(pr.Reviews),
			strings.Join(pr.HeadCommitBlockers, ", "),
			strings.Join(pr.HeadCommitApprovals, ", "),
		}
		retval += strings.Join(column, "|")
		retval += "\n"
	}
	return retval
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*pending_review.PullRequest) []*pending_review.PullRequestStatus {
	var out []*pending_review.PullRequestStatus
	for _, pr := range prs {
		if pr.GetDraft() {
			continue // Let's skip these
		}

		if len := len(pr.Labels); len > 0 {
			if len > 1 || !containsLabelNamed(pr.Labels, BUMP_VERSION) || !containsLabelNamed(pr.Labels, UNEXP_ERR) {
				continue // We know if there are labels then there's probably somethnig wrong!
			}
		}

		opt := &github.ListOptions{
			Page:    0,
			PerPage: 100,
		}
		for {
			review, resp, err := client.PullRequest.GatherRelevantReviews(context, "conan-io", "conan-center-index", pr, opt)
			if errors.Is(err, pending_review.ErrNoReviews) {
				break
			} else if err != nil {
				fmt.Printf("Problem getting list of reviews %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%+v\n", review)
			out = append(out, review)

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return out
}

func containsLabelNamed(slice []*github.Label, item string) bool {
	for _, a := range slice {
		if a.GetName() == item {
			return true
		}
	}
	return false
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func determineAndSetupCredentials(context context.Context) *http.Client {
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if exists {
		tokenService := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		return oauth2.NewClient(context, tokenService)
	} else {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("GitHub Username: ")
		username, _ := r.ReadString('\n')

		fmt.Print("GitHub Password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)

		tp := github.BasicAuthTransport{
			Username: strings.TrimSpace(username),
			Password: strings.TrimSpace(password),
		}

		return tp.Client()
	}
}
