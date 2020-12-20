package pending_review

import (
	"context"
	"net/http"

	"github.com/google/go-github/v33/github"
)

type Response = github.Response

type Client struct {
	*github.Client

	common service

	Repository *RepositoryService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	c := &Client{Client: github.NewClient(httpClient)}
	c.common.client = c
	c.Repository = (*RepositoryService)(&c.common)
	return c
}

type ListOptions = github.ListOptions

type RateLimit = github.Rate

func (c *Client) RateLimits(ctx context.Context) (*RateLimit, *Response, error) {
	rateLimit, resp, err := c.Client.RateLimits(ctx)
	if err != nil {
		return nil, nil, err
	}

	return rateLimit.Core, resp, nil
}
