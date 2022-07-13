package pending_review

import (
	"context"

	"github.com/google/go-github/v42/github"

	"gopkg.in/yaml.v3"
)

type ReviewerType string

const (
	Team      ReviewerType = "Team"
	Community ReviewerType = "Community"
)

type Reviewer struct {
	User      string       `yaml:"user"`
	Type      ReviewerType `yaml:"type"`
	Requested bool         `yaml:"request_reviews"`
}

type ConanCenterReviewers struct {
	Reviewers []Reviewer `yaml:"reviewers"`
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

	var reviewers ConanCenterReviewers
	err = yaml.Unmarshal([]byte(str), &reviewers)
	if err != nil {
		return nil, err
	}

	return &reviewers, nil
}
