package pending_review

import (
	"context"
	"errors"
	"time"

	"github.com/google/go-github/v42/github"
)

// RepositoryService handles communication with the repository related methods of the GitHub API
type RepositoryService service

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

// GetStatus fetches the complete status available for a ref
func (s *RepositoryService) GetStatus(ctx context.Context, owner string, repo string, ref string) (*CombinedStatus, *Response, error) {
	status, resp, err := s.client.Repositories.GetCombinedStatus(ctx, owner, repo, ref, &ListOptions{})
	if err != nil {
		return nil, resp, err
	}

	if status.GetTotalCount() == 0 {
		return nil, resp, ErrNoCommitStatus
	}

	return status, resp, nil
}
