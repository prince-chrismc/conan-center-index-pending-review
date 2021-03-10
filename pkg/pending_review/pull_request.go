package pending_review

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Category describing the type of change being introduced by the pull request
type Category int

// Category describing the type of change being introduced by the pull request
const (
	ADDED Category = iota
	EDIT  Category = iota
	BUMP  Category = iota
	DOCS  Category = iota
)

// ReviewSummary of a pull request based on the rules of conan-center-index.
// See https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md
// for more inforamtion
type ReviewSummary struct {
	Number              int
	OpenedBy            string
	Recipe              string
	Change              Category
	ReviewURL           string
	LastCommitSHA       string
	LastCommitAt        time.Time
	Reviews             int
	ValidApprovals      int
	IsMergeable         bool
	CciBotPassed        bool
	HeadCommitApprovals []string
	HeadCommitBlockers  []string
}

// ErrNoReviews indicated there were no reviews on a pull request and a summary could not be generated
var ErrNoReviews = errors.New("no reviews on pull request")

// PullRequestService handles communication with the pull request related methods of the GitHub API
type PullRequestService service

// GetReviewSummary of a specific pull request
func (s *PullRequestService) GetReviewSummary(ctx context.Context, owner string, repo string, pr *PullRequest) (*ReviewSummary, *Response, error) {
	p := &ReviewSummary{
		Number:        pr.GetNumber(),
		OpenedBy:      pr.GetUser().GetLogin(),
		ReviewURL:     pr.GetHTMLURL(),
		LastCommitSHA: pr.GetHead().GetSHA(),
	}

	diff, resp, err := s.determineTypeOfChange(ctx, owner, repo, p.Number, 10 /* recipes are currently 5-7 files */)
	if err != nil {
		return nil, resp, err
	}

	p.Recipe = diff.Recipe
	p.Change = diff.Change

	opt := &ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		reviews, resp, err := s.client.PullRequests.ListReviews(ctx, owner, repo, p.Number, opt)
		if err != nil {
			return nil, resp, err
		}

		if p.Reviews += len(reviews); p.Reviews < 1 { // Has not been looked at...
			date, _, err := s.client.Repository.GetCommitDate(ctx, pr.GetHead().GetRepo().GetOwner().GetLogin(), pr.GetHead().GetRepo().GetName(), p.LastCommitSHA)
			if err != nil {
				return nil, resp, err
			}
			p.LastCommitAt = date

			if isWithin24Hours(p.LastCommitAt) { // commited within 24hrs
				return p, resp, nil // let's save it!
			}

			return nil, resp, fmt.Errorf("%w", ErrNoReviews)
		}

		summary := ProcessReviewComments(reviews, p.LastCommitSHA)
		p.ValidApprovals += summary.ValidApprovals // FIXME: v2 refactor
		p.HeadCommitBlockers = append(p.HeadCommitBlockers, summary.HeadCommitBlockers...)
		p.HeadCommitApprovals = append(p.HeadCommitApprovals, summary.HeadCommitApprovals...)

		p.IsMergeable = summary.PriorityApproval && p.ValidApprovals >= 3 && len(p.HeadCommitBlockers) == 0

		status, _, err := s.client.Repository.GetCommitStatus(ctx, pr.GetBase().GetRepo().GetOwner().GetLogin(), pr.GetBase().GetRepo().GetName(), p.LastCommitSHA)
		if errors.Is(err, ErrNoCommitStatus) {
			p.CciBotPassed = false
		} else if err != nil {
			return nil, resp, err
		} else {
			p.CciBotPassed = status.GetState() == "success"
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	if len(p.HeadCommitApprovals) > 0 || p.Change == DOCS { // Always save documentation pull requests
		return p, resp, nil
	}

	return nil, resp, fmt.Errorf("%w", ErrNoReviews)
}

func isWithin24Hours(t time.Time) bool {
	return t.Add(time.Hour * 24).After(time.Now())
}

type change struct {
	Recipe string
	Change Category
}

// ErrInvalidChange in the commit files of the pull request which break the rules of CCI
var ErrInvalidChange = errors.New("the files, or lack thereof, make this PR invalid")

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

	change, err := getDiff(files[0])
	if err != nil {
		return nil, resp, err
	}

	for _, file := range files[1:] {
		obtained, err := getDiff(file)
		if err != nil {
			return nil, resp, err
		}

		if change.Recipe != obtained.Recipe { // PR should only be changing one recipe at a time
			return nil, resp, fmt.Errorf("%w", ErrInvalidChange)
		}

		if obtained.Change == EDIT {
			change.Change = EDIT // Any edit breaks the "new receipe" definition
		}
	}

	if len(files) == 2 {
		if strings.HasSuffix(files[0].GetFilename(), "conandata.yml") && strings.HasSuffix(files[1].GetFilename(), "config.yml") {
			change.Change = BUMP
		}
	}

	return change, resp, nil
}

// Expected format is
// - "recipes" , "<name>", "..."
// - "docs", "<filename>.md"
func getDiff(file *CommitFile) (*change, error) {
	segments := strings.SplitN(file.GetFilename(), "/", 3)
	if len(segments) < 2 { // Expected format is "recipes" , "<name>", "..."
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	folder := segments[0]
	title := segments[1]
	status := ADDED
	if file.GetStatus() != "added" {
		status = EDIT
	}
	if folder == "docs" {
		status = DOCS
		title = "docs"
	} else if folder != "recipes" {
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	return &change{title, status}, nil
}
