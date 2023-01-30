package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
	"golang.org/x/oauth2"
)

// ErrRateLimitReached by the client's login
var ErrRateLimitReached = errors.New("current login has exceed it's limit for requests that can be made")

func MakeClient(context context.Context, token string, target pending_review.WorkingRepository) (*pending_review.Client, error) {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

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
