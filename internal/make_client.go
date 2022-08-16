package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

// ErrRateLimitReached by the client's login
var ErrRateLimitReached = errors.New("current login has exceed it's limit for requests that can be made")

func MakeClient(context context.Context, c *cli.Context) (*pending_review.Client, error) {
	token := c.String("access-token")
	pem := c.String("app-pem")
	owner := c.String("repo-owner")
	repo := c.String("repo-name")
	target := pending_review.WorkingRepository{Owner: owner, Name: repo}

	var tokenService oauth2.TokenSource
	if pem == "" {
		tokenService = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	} else {
		source := GitHubInstallTokenSource{Owner: owner, Pem: pem}
		tokenService = oauth2.ReuseTokenSource(nil, &source)
	}

	client := pending_review.NewClient(oauth2.NewClient(context, tokenService), target)

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		return nil, err
	}

	// We have not exceeded the limit so we can continue
	fmt.Printf("Limit: %d \nRemaining: %d \n", rateLimit.Limit, rateLimit.Remaining)

	if rateLimit.Remaining <= 0 {
		return nil, ErrRateLimitReached
	}

	return client, nil
}

type GitHubAppAuthTokenSource struct {
	Pem string
}

func (s *GitHubAppAuthTokenSource) Token() (*oauth2.Token, error) {
	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.Pem))

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

	return &oauth2.Token{
		AccessToken: tokenString,
		Expiry:      exp,
	}, nil
}

type GitHubInstallTokenSource struct {
	Pem   string
	Owner string
}

func (s *GitHubInstallTokenSource) Token() (*oauth2.Token, error) {
	ctx := context.Background()

	tokenService := oauth2.ReuseTokenSource(nil, &GitHubAppAuthTokenSource{Pem: s.Pem})
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

	return &oauth2.Token{
		AccessToken: token.GetToken(),
		Expiry:      token.GetExpiresAt(),
	}, nil

}
