package pending_review

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Category describing the type of change being introduced by the pull request
type Category int

// Category describing the type of change being introduced by the pull request
const (
	NEW    Category = iota
	EDIT   Category = iota
	DOCS   Category = iota
	CONFIG Category = iota
)

// ReviewWeight attempts to qualify the amount of effort to review any given pull request
type ReviewWeight int

// ReviewWeight attempts to qualify the amount of effort to review any given pull request
const (
	TINY     ReviewWeight = iota
	SMALL    ReviewWeight = iota
	REGULAR  ReviewWeight = iota
	HEAVY    ReviewWeight = iota
	TOO_MUCH ReviewWeight = iota
)

// PullRequestSummary regarding its location in the review process of conan-center-index.
// See https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md
// for more information
type PullRequestSummary struct {
	Number        int
	OpenedBy      string
	CreatedAt     time.Time
	Recipe        string
	Change        Category
	Weight        ReviewWeight
	ReviewURL     string
	LastCommitSHA string
	LastCommitAt  time.Time
	CciBotPassed  bool
	Summary       Reviews
}

// ErrStoppedLabel indicates there is an issue with the pull request
var ErrStoppedLabel = errors.New("the pull request contains at least one label indicated that progress has stopped")

// ErrWorkRequired indicated there were no reviews on a pull request and a summary could not be generated
var ErrWorkRequired = errors.New("pull requests has comments which need to be addressed")

// ErrInvalidChange in the commit files of the pull request which break the rules of CCI
var ErrInvalidChange = errors.New("the files, or lack thereof, make this PR invalid")

// PullRequestService handles communication with the pull request related methods of the GitHub API
type PullRequestService service

// ListAllReviews lists all reviews on the specified pull request.
func (s *PullRequestService) ListAllReviews(ctx context.Context, owner string, repo string, number int) ([]*PullRequestReview, *Response, error) {
	var reviews []*PullRequestReview
	var resp *Response
	opt := &ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		newReviews, resp, err := s.client.PullRequests.ListReviews(ctx, owner, repo, number, opt)
		if err != nil {
			return nil, resp, err
		}

		reviews = append(reviews, newReviews...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return reviews, resp, nil
}

// GetReviewSummary of a specific pull request
func (s *PullRequestService) GetReviewSummary(ctx context.Context, owner string, repo string, reviewers *ConanCenterReviewers, pr *PullRequest) (*PullRequestSummary, *Response, error) {
	p := &PullRequestSummary{
		Number:        pr.GetNumber(),
		OpenedBy:      pr.GetUser().GetLogin(),
		CreatedAt:     pr.GetCreatedAt(),
		ReviewURL:     pr.GetHTMLURL(),
		LastCommitSHA: pr.GetHead().GetSHA(),
	}

	err := processLabels(pr.Labels)
	if err != nil {
		return nil, nil, err
	}

	diff, resp, err := s.determineTypeOfChange(ctx, owner, repo, p.Number, 14 /* recipes are currently 8-10 files */)
	if err != nil {
		return nil, resp, err
	}

	p.Recipe = diff.Recipe
	p.Change = diff.Change
	p.Weight = diff.Weight

	reviews, resp, err := s.ListAllReviews(ctx, owner, repo, p.Number)
	if err != nil {
		return nil, resp, err
	}

	reviews = FilterAuthor(reviews, p.OpenedBy)
	p.Summary = ProcessReviewComments(reviewers, reviews, p.LastCommitSHA)
	p.Summary.IsBump = false
	for _, label := range pr.Labels {
		switch label.GetName() {
		case "Bump version", "Bump dependencies":
			p.Summary.IsBump = true
		}
	}

	p.LastCommitAt, _, err = s.client.Repository.GetCommitDate(ctx, pr.GetHead().GetRepo().GetOwner().GetLogin(), pr.GetHead().GetRepo().GetName(), p.LastCommitSHA)
	if err != nil {
		return nil, resp, err
	}

	status, _, err := s.client.Repository.GetStatus(ctx, pr.GetBase().GetRepo().GetOwner().GetLogin(), pr.GetBase().GetRepo().GetName(), p.LastCommitSHA)
	if errors.Is(err, ErrNoCommitStatus) {
		p.CciBotPassed = false
	} else if err != nil {
		return nil, resp, err
	} else {
		p.CciBotPassed = status.GetState() == "success"
	}

	if bytes, err := json.Marshal(p); err == nil {
		fmt.Printf("%s\n", bytes)
	}
	err = evaluateSummary(p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

func (s *PullRequestService) determineTypeOfChange(ctx context.Context, owner string, repo string, number int, perPage int) (*change, *Response, error) {
	files, resp, err := s.client.PullRequests.ListFiles(ctx, owner, repo, number, &ListOptions{
		Page:    0,
		PerPage: perPage,
	})
	if err != nil {
		return nil, resp, err
	}

	if len(files) < 1 {
		return nil, resp, fmt.Errorf("%w", ErrInvalidChange)
	}

	change, err := processChangedFiles(files)
	if err != nil {
		return nil, resp, err
	}

	return change, resp, nil
}
