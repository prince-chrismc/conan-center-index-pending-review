package pending_review

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	// Review States
	CHANGE    = "CHANGES_REQUESTED"
	APPRVD    = "APPROVED"
	DISMISSED = "DISMISSED"

	// Author Associations
	COLLABORATOR = "COLLABORATOR"
	CONTRIBUTOR  = "CONTRIBUTOR"
	MEMBER       = "MEMBER"
)

type Status int

const (
	ADDED Status = iota
	EDIT  Status = iota
	BUMP  Status = iota
)

type PullRequestStatus struct {
	Number              int
	OpenedBy            string
	Recipe              string
	Change              Status
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

var ErrNoReviews = errors.New("no reviews on pull request")

type PullRequestService service

func (s *PullRequestService) GatherRelevantReviews(ctx context.Context, owner string, repo string, pr *PullRequest) (*PullRequestStatus, *Response, error) {
	p := &PullRequestStatus{
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

		atleastOneTeamApproval := false
		for _, review := range reviews {
			onBranchHead := p.LastCommitSHA == review.GetCommitID()
			reviewerName := review.GetUser().GetLogin()
			reviewerAssociation := review.GetAuthorAssociation()
			isC3iTeam := reviewerAssociation == MEMBER || reviewerAssociation == COLLABORATOR

			switch review.GetState() {
			case CHANGE:
				if isC3iTeam {
					p.HeadCommitBlockers = appendUnique(p.HeadCommitBlockers, reviewerName)
				}

				p.HeadCommitApprovals = removeUnique(p.HeadCommitApprovals, reviewerName)
			case APPRVD:
				if onBranchHead {
					p.HeadCommitApprovals = appendUnique(p.HeadCommitApprovals, reviewerName)
					if isC3iTeam {
						atleastOneTeamApproval = true
						p.ValidApprovals = p.ValidApprovals + 1
					}
				}

				switch reviewerName {
				case "madebr", "SpaceIm", "ericLemanissier", "prince-chrismc", "Croydon", "intelligide", "theirix", "gocarlos":
					p.ValidApprovals = p.ValidApprovals + 1
				default:
				}

				p.HeadCommitBlockers = removeUnique(p.HeadCommitBlockers, reviewerName)
			case DISMISSED:
				p.HeadCommitBlockers = removeUnique(p.HeadCommitBlockers, reviewerName)
				// Out-dated Approvals are transformed into comments https://github.com/conan-io/conan-center-index/pull/3855#issuecomment-770120073
			default:
			}
		}
		p.IsMergeable = atleastOneTeamApproval && p.ValidApprovals >= 3 && len(p.HeadCommitBlockers) == 0

		statuses, _, err := s.client.Repositories.ListStatuses(ctx, pr.GetBase().GetRepo().GetOwner().GetLogin(), pr.GetBase().GetRepo().GetName(), p.LastCommitSHA, &ListOptions{
			Page:    0,
			PerPage: 1,
		})
		if err != nil {
			return nil, resp, err
		}
		p.CciBotPassed = len(statuses) > 0 && statuses[0].GetState() == "success"

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	if len(p.HeadCommitApprovals) > 0 {
		return p, resp, nil
	}

	return nil, resp, fmt.Errorf("%w", ErrNoReviews)
}

func isWithin24Hours(t time.Time) bool {
	hours, _ := time.ParseDuration("24h")
	return t.Add(hours).After(time.Now())
}

func appendUnique(slice []string, elem string) []string {
	for _, e := range slice {
		if e == elem {
			return slice
		}
	}

	return append(slice, elem)
}

func removeUnique(slice []string, elem string) []string {
	for i, e := range slice {
		if e == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

type change struct {
	Recipe string
	Change Status
}

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

		if change.Recipe != obtained.Recipe { // PR shouls only be changing one recipe at a time
			return nil, resp, fmt.Errorf("%w", ErrInvalidChange)
		}

		if obtained.Change == EDIT {
			change.Change = EDIT // Any edit breaks the "new receipe"
		}
	}

	return change, resp, nil
}

func getDiff(file *CommitFile) (*change, error) {
	segments := strings.SplitN(file.GetFilename(), "/", 3)
	if len(segments) < 3 { // Expected format is "recipes" , "<name>", "..."
		return nil, fmt.Errorf("%w", ErrInvalidChange)
	}

	title := segments[1]
	status := ADDED
	if file.GetStatus() != "added" {
		status = EDIT
	}

	return &change{title, status}, nil
}
