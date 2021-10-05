package pending_review

import (
	"context"
	"net/http"

	"github.com/google/go-github/v39/github"
)

// Response is a GitHub API response.
type Response = github.Response

// PullRequest represents a GitHub pull request on a repository.
type PullRequest = github.PullRequest

// PullRequestReview represents a review of a pull request.
type PullRequestReview = github.PullRequestReview

// Label represents a GitHub label on an Issue
type Label = github.Label

// ListOptions specifies the optional parameters to various List methods that support offset pagination.
type ListOptions = github.ListOptions

// CommitFile represents a file modified in a commit.
type CommitFile = github.CommitFile

// RepoStatus represents the status of a repository at a particular reference.
type RepoStatus = github.RepoStatus

// CombinedStatus represents the combined status of a repository at a particular reference.
type CombinedStatus = github.CombinedStatus

// A Client manages communication with the GitHub API. This wraps the github.Client
// and provides convenient access to interupt information from CCI prespective
type Client struct {
	*github.Client

	common service

	Repository  *RepositoryService
	PullRequest *PullRequestService
}

type service struct {
	client *Client
}

// NewClient returns a new GitHub API client. Requires authentication.
func NewClient(httpClient *http.Client) *Client {
	c := &Client{Client: github.NewClient(httpClient)}
	c.common.client = c
	c.Repository = (*RepositoryService)(&c.common)
	c.PullRequest = (*PullRequestService)(&c.common)
	return c
}

// RateLimit represents the rate limit for the current client.
type RateLimit = github.Rate

// RateLimits returns the rate limits for the current client.
func (c *Client) RateLimits(ctx context.Context) (*RateLimit, *Response, error) {
	rateLimit, resp, err := c.Client.RateLimits(ctx)
	if err != nil {
		return nil, nil, err
	}

	return rateLimit.Core, resp, nil
}
