package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"golang.org/x/oauth2"
)

type GitHubAppAuthTokenSource struct {
}

func (s *GitHubAppAuthTokenSource) Token() (*oauth2.Token, error) {
	keyData, _ := os.ReadFile("conan-center-index-pending-review.2022-08-12.private-key.pem")
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	timeNow := time.Now()
	exp := timeNow.Add(time.Minute * 10) // 10 minutes in the future

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": 228148,                               // app_id
		"iat": timeNow.Add(time.Minute * -1).Unix(), // 1 minute ago
		"exp": exp.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	fmt.Println("Made token: ", tokenString)

	return &oauth2.Token{
		AccessToken: tokenString,
		Expiry:      exp,
	}, nil
}

type GitHubInstallTokenSource struct {
	Owner string
}

func (s *GitHubInstallTokenSource) Token() (*oauth2.Token, error) {
	ctx := context.Background()

	tokenService := oauth2.ReuseTokenSource(nil, &GitHubAppAuthTokenSource{})
	client := github.NewClient(oauth2.NewClient(ctx, tokenService))

	installs, _, err := client.Apps.ListInstallations(ctx, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	var installationId int64 = 0
	for _, install := range installs {
		if install.GetAccount().GetLogin() == s.Owner {
			installationId = install.GetID()
		}
	}

	if installationId == 0 {
		return nil, errors.New("unable to find installation for owner")
	}

	token, _, err := client.Apps.CreateInstallationToken(ctx, installationId, &github.InstallationTokenOptions{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Using installation's token: ", token.GetToken())

	return &oauth2.Token{
		AccessToken: token.GetToken(),
		Expiry:      token.GetExpiresAt(),
	}, nil

}

func main() {
	context := context.Background()
	source := GitHubInstallTokenSource{Owner: "prince-chrismc"}
	tokenService := oauth2.ReuseTokenSource(nil, &source)

	client := pending_review.NewClient(oauth2.NewClient(context, tokenService), pending_review.WorkingRepository{Owner: "prince-chrismc", Name: "conan-center-index-pending-review"})

	rateLimit, _, _ := client.RateLimits(context)

	fmt.Println(nil, "has", rateLimit.Remaining, "requests left")

	_, _, err := client.Issues.CreateComment(context, "prince-chrismc", "conan-center-index-pending-review", 85, &github.IssueComment{
		Body: github.String("# Hello World!\n\n Trying to make a bot for https://github.com/prince-chrismc/conan-center-index-pending-review/issues/1"),
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
