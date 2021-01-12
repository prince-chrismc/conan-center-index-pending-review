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
	BUMP_VERSION = "Bump version"
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

	// https://github.com/prince-chrismc/conan-center-index-pending-review/issues/5#issuecomment-754112342
	isDifferent, err := validateContentIsDifferent(context, client, string(bytes))
	if err != nil {
		fmt.Printf("Problem getting original issue content %v\n", err)
		os.Exit(1)
	}
	if !isDifferent {
		fmt.Println("the obtained content is identical to the new result.")
		return // The published results are the same no need to update the table.
	}

	_, _, err = client.Issues.Edit(context, "prince-chrismc", "conan-center-index-pending-review", 1, &github.IssueRequest{
		Body: github.String(`## :sparkles: Pull Requests Pending Review Summary!

### :ballot_box_with_check: Selection Criteria:

- No reviews and commited to in the last 24hrs
- There has been at least one approval

### :nerd_face: Please Review!

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + formatPullRequestToMarkdownRows(retval, false) + `

### :heavy_check_mark: Ready to Merge

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + formatPullRequestToMarkdownRows(retval, true) + "\n\n<details><summary>Raw JSON data</summary>\n\n```json\n" + string(bytes) + "\n```\n\n</details>"),
	})
	if err != nil {
		fmt.Printf("Problem editing issue %v\n", err)
		os.Exit(1)
	}
}

func formatPullRequestToMarkdownRows(prs []*pending_review.PullRequestStatus, canMerge bool) string {
	var retval string
	for _, pr := range prs {
		if pr.IsMergeable == canMerge {
			continue
		}

		title := "recipe"
		switch pr.Change {
		case pending_review.ADDED:
			title = ":new: " + pr.Recipe
			break
		case pending_review.EDIT:
			title = ":memo: " + pr.Recipe
			break
		case pending_review.BUMP:
			title = ":arrow_up: " + pr.Recipe
			break
		}

		columns := []string{
			fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
			fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
			title,
			fmt.Sprint(pr.Reviews),
			strings.Join(pr.HeadCommitBlockers, ", "),
			strings.Join(pr.HeadCommitApprovals, ", "),
		}
		retval += strings.Join(columns, "|")
		retval += "\n"
	}
	return retval
}

func validateContentIsDifferent(context context.Context, client *pending_review.Client, expected string) (bool, error) {
	issue, _, err := client.Issues.Get(context, "prince-chrismc", "conan-center-index-pending-review", 1)
	if err != nil {
		return false, err
	}
	content := issue.GetBody()

	rawJsonStart := strings.Index(content, "```json\n")
	rawJsonEnd := strings.Index(content, "\n```\n")

	if rawJsonStart == -1 || rawJsonEnd == -1 {
		return false, errors.New("content did not contain the expected raw JSON section")
	}

	obtained := content[rawJsonStart+len("```json\n") : rawJsonEnd]

	return obtained != expected, nil
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*pending_review.PullRequest) []*pending_review.PullRequestStatus {
	var out []*pending_review.PullRequestStatus
	for _, pr := range prs {
		if pr.GetDraft() {
			continue // Let's skip these
		}

		isBump := false
		if len := len(pr.Labels); len > 0 {
			for _, label := range pr.Labels {
				name := label.GetName()
				if name == BUMP_VERSION {
					isBump = true
				}
			}

			if !isBump {
				continue // We know if there are certain labels then there's probably something wrong!
			}
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

		fmt.Printf("%+v\n", review)
		out = append(out, review)
	}
	return out
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
