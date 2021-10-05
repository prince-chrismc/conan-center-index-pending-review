package pending_review

import (
	"context"
	"errors"
	"time"

	"github.com/google/go-github/v39/github"
)

// RepositoryService handles communication with the repository related methods of the GitHub API
type RepositoryService service

// RepositorySumarry provides a basic overview of a specific repository
type RepositorySumarry struct {
	Name            string
	Owner           string
	FullName        string
	Description     string
	StarsCount      int
	ForksCount      int
	OpenIssuesCount int
}

// GetSummary of a specific repository on GitHub
func (s *RepositoryService) GetSummary(ctx context.Context, owner string, repo string) (*RepositorySumarry, *Response, error) {
	repos, resp, err := s.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, resp, err
	}

	return &RepositorySumarry{
		Name:            repos.GetName(),
		Owner:           repos.GetOwner().GetLogin(),
		FullName:        repos.GetFullName(),
		Description:     repos.GetDescription(),
		ForksCount:      repos.GetForksCount(),
		StarsCount:      repos.GetStargazersCount(),
		OpenIssuesCount: repos.GetOpenIssuesCount(),
	}, resp, nil
}

// GetCommitDate fetches the specified commit's authorship date
func (s *RepositoryService) GetCommitDate(ctx context.Context, owner string, repo string, sha string) (time.Time, *Response, error) {
	commit, resp, err := s.client.Repositories.GetCommit(ctx, owner, repo, sha, &github.ListOptions{})
	if err != nil {
		return time.Time{}, resp, err
	}
	return commit.GetCommit().GetAuthor().GetDate(), resp, nil
}

// ErrNoCommitStatus available
var ErrNoCommitStatus = errors.New("no repository status avialble for this commit")

// GetCommitStatus fetches the latest/most recent status available for a commit
func (s *RepositoryService) GetCommitStatus(ctx context.Context, owner string, repo string, sha string) (*RepoStatus, *Response, error) {
	statuses, resp, err := s.client.Repositories.ListStatuses(ctx, owner, repo, sha, &ListOptions{
		Page:    0,
		PerPage: 1,
	})
	if err != nil {
		return nil, resp, err
	}

	if len(statuses) == 0 {
		return nil, resp, ErrNoCommitStatus
	}

	return statuses[0], resp, nil
}
