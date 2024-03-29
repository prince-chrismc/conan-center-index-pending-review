package pending_review

import (
	"context"

	"github.com/google/go-github/v45/github"

	"gopkg.in/yaml.v3"
)

type ReviewerType string

const (
	Team      ReviewerType = "team"
	Community ReviewerType = "community"
	Unofficial ReviewerType = "unofficial" // This is an add-in for pending_review to track extra approvals
)

type Reviewer struct {
	User      string       `yaml:"user"`
	Type      ReviewerType `yaml:"type"`
	Requested bool         `yaml:"request_reviews"`
}

type ConanCenterReviewers struct {
	Reviewers []Reviewer `yaml:"reviewers"`
}

func (reviewers *ConanCenterReviewers) IsTeamMember(reviewerName string) bool {
	for _, v := range reviewers.Reviewers {
		if v.User == reviewerName && v.Type == Team {
			return true
		}
	}

	return false
}

func (reviewers *ConanCenterReviewers) IsCommunityMember(reviewerName string) bool {
	for _, v := range reviewers.Reviewers {
		if v.User == reviewerName && v.Type == Community {
			return true
		}
	}

	return false
}

func parseReviewers(str string) (*ConanCenterReviewers, error) {
	var reviewers ConanCenterReviewers
	err := yaml.Unmarshal([]byte(str), &reviewers)
	if err != nil {
		return nil, err
	}

	return &reviewers, nil
}

func DownloadKnownReviewersList(context context.Context, client *Client) (*ConanCenterReviewers, error) {
	fileContent, _, _, err := client.Repositories.GetContents(context, "conan-io", "conan-center-index", ".c3i/reviewers.yml",
		&github.RepositoryContentGetOptions{Ref: "master"})
	if err != nil {
		return nil, err
	}

	str, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	reviewers, err := parseReviewers(str)
	if err != nil {
		return nil, err
	}

	return reviewers, nil
}
