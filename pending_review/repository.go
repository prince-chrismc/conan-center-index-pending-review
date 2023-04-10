package pending_review

import (
	"context"
	"errors"
	"time"

	"github.com/google/go-github/v45/github"
)

// RepositoryService handles communication with the repository related methods of the GitHub API
type RepositoryService service

// GetCommitDate fetches the specified commit's authorship date
func (s *RepositoryService) GetCommitDate(ctx context.Context, owner string, repo string, sha string) (time.Time, *Response, error) {
	commit, resp, err := s.client.Repositories.GetCommit(ctx, owner, repo, sha, &github.ListOptions{})
	if err != nil {
		return time.Time{}, resp, err
	}

	// Here we need to get the committer date since this appears to be when the commit was pushed to the server rather
	// then when it was initial created, for examples https://api.github.com/repos/conan-io/conan-center-index/git/commits/6b173fd061c77e5eb51990f372d9c138f14bd7fa
	// Where the OP force pushed the same commit for 5 months.
	return commit.GetCommit().GetCommitter().GetDate(), resp, nil
}

// ErrNoCommitStatus available
var ErrNoCommitStatus = errors.New("no repository status available for this commit")

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
