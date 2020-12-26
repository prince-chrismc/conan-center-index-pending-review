package pending_review

import (
	"context"
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
