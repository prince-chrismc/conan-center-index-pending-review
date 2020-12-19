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
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

// Model
type Package struct {
	FullName        string
	Description     string
	StarsCount      int
	ForksCount      int
	LastUpdatedBy   string
	OpenIssuesCount int
}

type PullRequest struct {
	Number    int
	Comments  int
	ReviewUrl string
}

func main() {
	context := context.Background()

	var httpClient *http.Client

	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if exists {
		tokenService := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(context, tokenService)
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

		httpClient = tp.Client()
	}

	client := github.NewClient(httpClient)

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		fmt.Printf("Problem in getting rate limit information %v\n", err)
		return
	}

	fmt.Printf("Limit: %d \nRemaining %d \n", rateLimit.Core.Limit, rateLimit.Core.Remaining)

	repo, _, err := client.Repositories.Get(context, "conan-io", "conan-center-index")
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName:        *repo.FullName,
		Description:     *repo.Description,
		ForksCount:      *repo.ForksCount,
		StarsCount:      *repo.StargazersCount,
		OpenIssuesCount: *repo.OpenIssuesCount,
	}

	fmt.Printf("%+v\n", pack)

	pulls, _, err := client.PullRequests.List(context, "conan-io", "conan-center-index", &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 3,
		},
	})
	for _, pr := range pulls {
		p := PullRequest{Number: pr.GetNumber(), Comments: pr.GetComments(), ReviewUrl: pr.GetReviewCommentsURL()}
		fmt.Printf("%+v\n", p)

		reviews, _, err := client.PullRequests.ListReviews(context, "conan-io", "conan-center-index", p.Number, &github.ListOptions{
			Page:    0,
			PerPage: 10,
		})
		if err != nil {
			fmt.Printf("Problem getting reviews information %v\n", err)
			os.Exit(1)
		}

		for _, review := range reviews {
			fmt.Printf("%s (%s): '%s' on commit %s\n", review.User.GetLogin(), review.GetAuthorAssociation(), review.GetState(), review.GetCommitID())
		}
	}
}

