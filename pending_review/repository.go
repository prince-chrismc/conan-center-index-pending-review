package pending_review

import (
	"context"
	"errors"
	"time"
)

type RepositoryService service

type Repository struct {
	Name            string
	Owner           string
	FullName        string
	Description     string
	StarsCount      int
	ForksCount      int
	OpenIssuesCount int
}

func (s *RepositoryService) Get(ctx context.Context, owner string, repo string) (*Repository, *Response, error) {
	repos, resp, err := s.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, resp, err
	}

	return &Repository{
		Name:            repos.GetName(),
		Owner:           repos.GetOwner().GetLogin(),
		FullName:        repos.GetFullName(),
		Description:     repos.GetDescription(),
		ForksCount:      repos.GetForksCount(),
		StarsCount:      repos.GetStargazersCount(),
		OpenIssuesCount: repos.GetOpenIssuesCount(),
	}, resp, nil
}

func (s *RepositoryService) GetCommitDate(ctx context.Context, owner string, repo string, sha string) (time.Time, *Response, error) {
	commit, resp, err := s.client.Repositories.GetCommit(ctx, owner, repo, sha)
	if err != nil {
		return time.Time{}, resp, err
	}
	return commit.GetCommit().GetAuthor().GetDate(), resp, nil
}

var ErrNoCommitStatus = errors.New("no repository status avialble for this commit")

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
