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
	"time"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v1/pending_review"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

const (
	BUMP = "Bump version"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65#gistcomment-2557682
func duration(d time.Duration) string {
	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd%s", days, d)

	return b.String()
}

type stats struct {
	Open    int
	Draft   int
	Review  int
	Merge   int
	Stale   int
	Failed  int
	Blocked int
	Age     time.Duration
}

func main() {
	context := context.Background()
	client := pending_review.NewClient(determineAndSetupCredentials(context))

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

	var stats stats
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
		stats.Age = time.Duration(((stats.Age.Nanoseconds() * int64(stats.Open)) + (s.Age.Nanoseconds() * int64(s.Open))) / int64(stats.Open+s.Open))
		stats.Open += s.Open
		stats.Draft += s.Draft
		stats.Stale += s.Stale
		stats.Failed += s.Failed
		stats.Blocked += s.Blocked
		stats.Merge += s.Merge
		stats.Review += s.Review

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
		return // The published results are the same, no need to update the table.
	}

	var ready string
	if stats.Merge > 0 {
		ready = `

### :heavy_check_mark: Ready to Merge (` + fmt.Sprint(stats.Merge) + `)

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + formatPullRequestToMarkdownRows(retval, true)
	}

	_, _, err = client.Issues.Edit(context, "prince-chrismc", "conan-center-index-pending-review", 1, &github.IssueRequest{
		Body: github.String(`## :sparkles: Pull Requests Pending Review Summary!

### :ballot_box_with_check: Selection Criteria:

- There has been at least one approval (at any point)
- No reviews and commited to in the last 24hrs
- No labels with exception to "bump version" and "docs"

### :nerd_face: Please Review! 

There are **` + fmt.Sprint(stats.Review) + `** pull requests currently under way :eyes:

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + formatPullRequestToMarkdownRows(retval, false) + ready + `

#### :bar_chart: Statistics

> :warning: These are just rough metrics counthing the labels and may not reflect the acutal state of pull requests

- Commit: ` + os.Getenv("GITHUB_SHA") + `
- PRs
   - Open: ` + fmt.Sprint(stats.Open) + `
   - Draft: ` + fmt.Sprint(stats.Draft) + `
   - Age: ` + duration(stats.Age) + `
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
}

func formatPullRequestToMarkdownRows(prs []*pending_review.PullRequestStatus, canMerge bool) string {
	var retval string
	for _, pr := range prs {
		if pr.IsMergeable != canMerge {
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
		case pending_review.DOCS:
			title = ":green_book: " + pr.Recipe
			break
		}

		if !pr.CciBotPassed && pr.IsMergeable {
			title = ":warning: " + pr.Recipe
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

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*pending_review.PullRequest) ([]*pending_review.PullRequestStatus, stats) {
	var stats stats
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
		if errors.Is(err, pending_review.ErrNoReviews) && isDoc {
			err = nil  // Always save documentation pull requests
		} else if errors.Is(err, pending_review.ErrNoReviews) || errors.Is(err, pending_review.ErrInvalidChange) {
			continue
		} else if err != nil {
			fmt.Printf("Problem getting list of reviews %v\n", err)
			os.Exit(1)
		}

		if isBump {
			review.Change = pending_review.BUMP // FIXME: It would be nice for this logic to be internal
		}

		stats.Review++
		if review.IsMergeable {
			stats.Merge++
		}

		fmt.Printf("%+v\n", review)
		out = append(out, review)
	}
	return out, stats
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
