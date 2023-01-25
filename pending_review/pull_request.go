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

// ErrStoppedLabel indicates the pull request only has minor impact and is automatically handled by the bot, does not require attention
var ErrBumpLabel = errors.New("the pull request is labelled as bump and will automatically be merged")

// ErrNoReviews indicated there were no reviews on a pull request and a summary could not be generated
var ErrNoReviews = errors.New("no reviews on pull request")

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

	if p.Change == DOCS { // Always save documentation pull requests
		return p, resp, nil
	}

	if p.Change == CONFIG && p.CciBotPassed { // Always save infrastructure pull requests that are passing
		return p, resp, nil
	}

	if p.Summary.Count < 1 { // Has not been looked at...
		return p, resp, nil // let's save it! So it can get some attention
	}

	if len(p.Summary.Approvals) > 0 { // It's been approved!
		return p, resp, nil
	}

	if p.LastCommitAt.After(p.Summary.LastReview.SubmittedAt) { // OP has presumably applied review comments
		return p, resp, nil // Let's save it so it gets another pass
	}

	return nil, resp, fmt.Errorf("%w", ErrNoReviews)
}

func processLabels(labels []*Label) error {
	for _, label := range labels {
		switch label.GetName() {
		case "bug", "stale", "Failed", "infrastructure", "blocked", "duplicate", "invalid":
			return fmt.Errorf("%w", ErrStoppedLabel)
		case "Bump version", "Bump dependencies":
			return fmt.Errorf("%w", ErrBumpLabel)
		default:
			continue
		}
	}

	return nil
}

type change struct {
	Recipe string
	Change Category
	Weight ReviewWeight
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

func processChangedFiles(files []*CommitFile) (*change, error) {
	if len(files) < 1 {
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	change, err := getDiff(files[0])
	if err != nil {
		return nil, err
	}

	addition := files[0].GetAdditions()
	deletions := files[0].GetDeletions()
	for _, file := range files[1:] {
		obtained, err := getDiff(file)
		if err != nil {
			return nil, err
		}

		if change.Recipe != obtained.Recipe {
			return nil, fmt.Errorf("%w", ErrInvalidChange)
		}

		if change.Change == NEW && obtained.Change == EDIT {
			change.Change = EDIT
		}

		addition += file.GetAdditions()
		deletions += file.GetDeletions()
	}

	//if len(files) <= 2 && addition <= 10 && deletions == 0 {
	if len(files) <= 2 && (addition+deletions) <= 10 {
		change.Weight = TINY
	} else if len(files) <= 4 && (addition+deletions) <= 25 {
		change.Weight = SMALL
	} else if len(files) <= 5 && (addition+deletions) <= 100 {
		change.Weight = REGULAR
	} else if len(files) <= 9 && addition <= 225 && deletions == 0 { // Basic new recipe addition with `test_v1_package`
		change.Weight = REGULAR
	} else if len(files) > 12 || (addition+deletions) >= 500 {
		change.Weight = TOO_MUCH
	}

	return change, nil
}

func getDiff(file *CommitFile) (*change, error) {
	// Expected format is: "folder" , "<name>", "..."
	// Other changes are 3-9 months so not worth supporting
	segments := strings.SplitN(file.GetFilename(), "/", 3)
	if len(segments) < 2 {
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	folder := segments[0]
	title := segments[1]
	status := NEW
	if file.GetStatus() != "added" {
		status = EDIT
	}

	switch folder {
	case "docs":
		status = DOCS
		title = "docs"
	case ".github":
		status = CONFIG
		title = ".github"
	case ".c3i":
		status = CONFIG
		title = ".c3i"
	case "linter":
		status = CONFIG
		title = "linter"
	case "recipes":
	default:
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	return &change{title, status, HEAVY}, nil // Default to heavy to make the calculation easier in `processChangedFiles`
}
