package pending_review

import (
	"context"
	"fmt"
	"strings"
)

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

	if len(files) <= 2 && (addition+deletions) <= 10 {
		change.Weight = TINY
	} else if len(files) <= 4 && (addition+deletions) <= 40 {
		change.Weight = SMALL
	} else if len(files) <= 6 && (addition+deletions) <= 30 { // More files but less LOC
		change.Weight = SMALL
	} else if len(files) <= 7 && (addition+deletions) <= 100 {
		change.Weight = REGULAR
	} else if len(files) <= 9 && addition <= 225 && deletions == 0 { // Basic new recipe addition with `test_v1_package`
		change.Weight = REGULAR
	} else if len(files) <= 12 || (addition+deletions) < 500 {
		change.Weight = HEAVY
	} else {
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
