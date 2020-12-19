package main

import (
	"bufio"
	"context"
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

// Model
type Package struct {
	FullName        string
	Description     string
	StarsCount      int
	ForksCount      int
	OpenIssuesCount int
}

type PullRequest struct {
	Number        int
	Reviews       int
	LastCommitSHA string
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

	repo, _, err := client.Repositories.Get(context, "conan-io", "conan-center-index")
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n-----\n", Package{
		FullName:        *repo.FullName,
		Description:     *repo.Description,
		ForksCount:      *repo.ForksCount,
		StarsCount:      *repo.StargazersCount,
		OpenIssuesCount: *repo.OpenIssuesCount,
	})

	pulls, _, err := client.PullRequests.List(context, "conan-io", "conan-center-index", &github.PullRequestListOptions{
		// Sort:      "long-running",
		// Direction: "desc",
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	})

	for _, pr := range pulls {
		p := PullRequest{
			Number: pr.GetNumber(),
		}

		if len(pr.Labels) > 0 {
			fmt.Printf("#%d has labels: ", p.Number)
			for _, label := range pr.Labels {
				fmt.Printf(" %s", label.GetName())
			}
			fmt.Println()
			continue // We know if there are labels then there's probably somethnig wrong!
		}

		reviews, _, err := client.PullRequests.ListReviews(context, "conan-io", "conan-center-index", p.Number, &github.ListOptions{
			Page:    0,
			PerPage: 100,
		})
		if err != nil {
			fmt.Printf("Problem getting reviews information %v\n", err)
			os.Exit(1)
		}
		p.Reviews = len(reviews)

		if p.Reviews < 1 {
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
		fmt.Printf("%+v\n", p)

		for _, review := range reviews {
			if review.GetState() != "APPROVED" {
				continue // Let's ignore the rest!
			}
			fmt.Printf("%s (%s): '%s' on commit %s\n", review.User.GetLogin(), review.GetAuthorAssociation(), review.GetState(), review.GetCommitID())
		}
	}

	issueComment, _, err := client.Issues.GetComment(context, "conan-io", "conan-center-index", 771457642)
	if err != nil {
		fmt.Printf("Problem getting issue comment %v\n", err)
		os.Exit(1)
}

	issueComment.Body = github.String("Hello World, From Action!")
	issueComment, _, err = client.Issues.EditComment(context, "conan-io", "conan-center-index", 771457642, issueComment)
	if err != nil {
		fmt.Printf("Problem editing issue comment %v\n", err)
		os.Exit(1)
	}
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
