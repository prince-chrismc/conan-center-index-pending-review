package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/pending_review"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

const (
	// Labels
	BUMP_VERSION = "Bump Version"

	// Review States
	CHANGE = "CHANGES_REQUESTED"
	APPRVD = "APPROVED"

	// Author Associations
	COLLABORATOR = "COLLABORATOR"
	CONTRIBUTOR  = "CONTRIBUTOR"
	MEMBER       = "MEMBER"
)

type PullRequest struct {
	Number              int
	OpenedBy            string
	ReviewURL           string
	LastCommitSHA       string
	Reviews             int
	AtLeastOneApproval  bool
	HeadCommitApprovals []string
	HeadCommitBlockers  []string
}

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

	var retval []PullRequest
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
		Body: github.String("Hello World, From Action!\n\n```json\n" + string(bytes) + "\n```"),
	})
	if err != nil {
		fmt.Printf("Problem editing issue %v\n", err)
		os.Exit(1)
	}
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, prs []*github.PullRequest) []PullRequest {
	var out []PullRequest
	for _, pr := range prs {
		if pr.GetDraft() {
			continue // Let's skip these
		}

		if len := len(pr.Labels); len > 0 {
			if len > 1 || !containsLabelNamed(pr.Labels, BUMP_VERSION) {
				continue // We know if there are labels then there's probably somethnig wrong!
			}
		}

		p := PullRequest{
			Number:        pr.GetNumber(),
			OpenedBy:      pr.GetUser().GetLogin(),
			ReviewURL:     pr.GetHTMLURL(),
			LastCommitSHA: pr.GetHead().GetSHA(),
		}

		opt := &github.ListOptions{
			Page:    0,
			PerPage: 100,
		}
		for {
			reviews, resp, err := client.PullRequests.ListReviews(context, "conan-io", "conan-center-index", p.Number, opt)
			if err != nil {
				fmt.Printf("Problem getting list of reviews %v\n", err)
				os.Exit(1)
			}

			if p.Reviews = len(reviews); p.Reviews < 1 {
				out = append(out, p)
				continue // Has not been looked at, let's save it!
			}

			for _, review := range reviews {
				onBranchHead := p.LastCommitSHA == review.GetCommitID()
				reviewerName := review.User.GetLogin()
				reviewerAssociation := review.GetAuthorAssociation()
				isC3iTeam := reviewerAssociation == MEMBER || reviewerAssociation == COLLABORATOR

				switch state := review.GetState(); state {
				case CHANGE:
					fmt.Printf("%s (%s): '%s' on commit %s\n", reviewerName, reviewerAssociation, state, review.GetCommitID())
					if onBranchHead && isC3iTeam {
						p.HeadCommitBlockers = append(p.HeadCommitBlockers, reviewerName)
					}
				case APPRVD:
					p.AtLeastOneApproval = true
					fmt.Printf("%s (%s): '%s' on commit %s\n", reviewerName, reviewerAssociation, state, review.GetCommitID())

					if onBranchHead {
						p.HeadCommitApprovals = append(p.HeadCommitApprovals, reviewerName)
					}
				default:
				}
			}

			if p.AtLeastOneApproval {
				out = append(out, p)
			}

			fmt.Printf("%+v\n", p)

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
