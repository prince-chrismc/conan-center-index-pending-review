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

type PullRequest struct {
	Number              int
	Reviews             int
	LastCommitSHA       string
	AtLeastOneApproval  bool
	HeadCommitApprovals int
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

	pulls, _, err := client.PullRequests.List(context, "conan-io", "conan-center-index", &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	})

	var retval []PullRequest
	for _, pr := range pulls {
		p := PullRequest{
			Number: pr.GetNumber(),
		}

		if len := len(pr.Labels); len > 0 {
			if len > 1 || !containsLabelNamed(pr.Labels, "Bump Version") {
			continue // We know if there are labels then there's probably somethnig wrong!
		}
		}

		reviews, _, err := client.PullRequests.ListReviews(context, "conan-io", "conan-center-index", p.Number, &github.ListOptions{
			Page:    0,
			PerPage: 100,
		})
		if err != nil {
			fmt.Printf("Problem getting reviews information %v\n", err)
			os.Exit(1)
		}

		if p.Reviews = len(reviews); p.Reviews < 1 {
			continue // Has not been looked at, let's skip!
		}

		commits, _, err := client.PullRequests.ListCommits(context, "conan-io", "conan-center-index", p.Number, &github.ListOptions{
			Page:    0,
			PerPage: 100,
		})
		if err != nil {
			fmt.Printf("Problem getting reviews information %v\n", err)
			os.Exit(1)
		}

		head := commits[len(commits)-1]
		p.LastCommitSHA = head.GetSHA()

		for _, review := range reviews {
			if review.GetState() != "APPROVED" {
				continue // Let's ignore the rest!
			}

			p.AtLeastOneApproval = true
			fmt.Printf("%s (%s): '%s' on commit %s\n", review.User.GetLogin(), review.GetAuthorAssociation(), review.GetState(), review.GetCommitID())

			if p.LastCommitSHA == review.GetCommitID() {
				p.HeadCommitApprovals = p.HeadCommitApprovals + 1
		}
		}

		if p.AtLeastOneApproval {
			retval = append(retval, p)
		}

		fmt.Printf("%+v\n", p)
	}

	bytes, err := json.MarshalIndent(retval, "", " ")
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

func containsLabelNamed(slice []*github.Label, item string) bool {
	for _, a := range slice {
		if a.GetName() == item {
			return true
		}
	}
	return false
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
