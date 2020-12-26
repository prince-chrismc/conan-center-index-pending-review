package pending_review

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	// Review States
	CHANGE = "CHANGES_REQUESTED"
	APPRVD = "APPROVED"

	// Author Associations
	COLLABORATOR = "COLLABORATOR"
	CONTRIBUTOR  = "CONTRIBUTOR"
	MEMBER       = "MEMBER"
)

type PullRequestService service

type PullRequestStatus struct {
	Number              int
	OpenedBy            string
	ReviewURL           string
	LastCommitSHA       string
	Reviews             int
	AtLeastOneApproval  bool
	HeadCommitApprovals []string
	HeadCommitBlockers  []string
}

var ErrNoReviews = errors.New("no reviews on pull request")

func (s *PullRequestService) GatherRelevantReviews(ctx context.Context, owner string, repo string, pr *PullRequest, opts *ListOptions) (*PullRequestStatus, *Response, error) {
	p := &PullRequestStatus{
		Number:        pr.GetNumber(),
		OpenedBy:      pr.GetUser().GetLogin(),
		ReviewURL:     pr.GetHTMLURL(),
		LastCommitSHA: pr.GetHead().GetSHA(),
	}

	reviews, resp, err := s.client.PullRequests.ListReviews(ctx, owner, repo, p.Number, opts)
	if err != nil {
		return nil, resp, err
	}

	if p.Reviews = len(reviews); p.Reviews < 1 { // Has not been looked at...
		hours, _ := time.ParseDuration("24h")
		if pr.GetCreatedAt().Add(hours).After(time.Now()) { // created within 24hrs
			return p, resp, nil // let's save it!
		}
		return nil, resp, fmt.Errorf("%w", ErrNoReviews)
	}

	for _, review := range reviews {
		onBranchHead := p.LastCommitSHA == review.GetCommitID()
		reviewerName := review.User.GetLogin()
		reviewerAssociation := review.GetAuthorAssociation()
		isC3iTeam := reviewerAssociation == MEMBER || reviewerAssociation == COLLABORATOR

		switch state := review.GetState(); state {
		case CHANGE:
			if onBranchHead && isC3iTeam {
				p.HeadCommitBlockers = append(p.HeadCommitBlockers, reviewerName)
			}
		case APPRVD:
			p.AtLeastOneApproval = true
			if onBranchHead {
				p.HeadCommitApprovals = append(p.HeadCommitApprovals, reviewerName)
			}
		default:
		}
	}

	if p.AtLeastOneApproval {
		return p, resp, nil
	}

	return nil, resp, fmt.Errorf("%w", ErrNoReviews)
}
