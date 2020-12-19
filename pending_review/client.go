package pending_review

import (
	"context"
	"net/http"

	"github.com/google/go-github/v33/github"
)

type Client struct {
	*github.Client
}

func NewClient(httpClient *http.Client) *Client {
	c := &Client{}
	c.Client = github.NewClient(httpClient)
	return c
}

type ListOptions = github.ListOptions

type RateLimit = github.Rate

func (c *Client) RateLimits(ctx context.Context) (*RateLimit, *github.Response, error) {
	rateLimit, resp, err := c.Client.RateLimits(ctx)
	if err != nil {
		return nil, nil, err
	}

	return rateLimit.Core, resp, nil
}
