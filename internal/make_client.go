package internal

import (
	"context"
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"golang.org/x/oauth2"
)

func MakeClient(context context.Context, token string) (*pending_review.Client, error) {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	client := pending_review.NewClient(oauth2.NewClient(context, tokenService))

	// Get Rate limit information
	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		return nil, err
	}

	// We have not exceeded the limit so we can continue
	fmt.Printf("Limit: %d \nRemaining: %d \n", rateLimit.Limit, rateLimit.Remaining)

	return client, nil
}
