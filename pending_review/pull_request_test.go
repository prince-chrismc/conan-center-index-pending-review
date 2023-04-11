package pending_review

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestGetReviewSummary16144(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/pulls/16144/reviews").
		Reply(200).
		File("test-data/16144/reviews.json")
	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/pulls/16144/files").
		Reply(200).
		File("test-data/16144/files.json")
	gock.New("https://api.github.com").
		Get("/repos/SpaceIm/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba").
		Reply(200).
		File("test-data/16144/e2aa65c961d48d688dd5450811229eb1d62649ba.json")
	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba/status").
		Reply(200).
		File("test-data/16144/status.json")

	client := NewClient(&http.Client{}, WorkingRepository{
		Owner: "prince-chrismc", Name: "conan-center-index-pending-review",
	})
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{
		{User: "czoido", Type: "team", Requested: false},
		{User: "memsharded", Type: "team", Requested: false},
		{User: "uilianries", Type: "team", Requested: true},
		{User: "SpaceIm", Type: "community", Requested: false},
		{User: "ericLemanissier", Type: "community", Requested: false},
		{User: "prince-chrismc", Type: "team", Requested: false},
		{User: "Croydon", Type: "community", Requested: false},
		{User: "toge", Type: "community", Requested: false},
	}}

	body, err := os.ReadFile("test-data/16144/pull-request.json")
	assert.Equal(t, nil, err)

	pr := parsePullRequestJSON(t, string(body))

	review, _, err := client.PullRequest.GetReviewSummary(context.Background(), "conan-io", "conan-center-index", &reviewers, pr)
	assert.Equal(t, nil, err)

	const layout = "2006-01-02 15:04:05 -0700 MST"
	createdAt, err := time.Parse(layout, "2023-02-19 15:10:36 +0000 UTC") // This is the debug time from `%+v` formatter
	assert.Equal(t, nil, err)
	lastCommitAt, err := time.Parse(layout, "2023-02-19 15:10:08 +0000 UTC")
	assert.Equal(t, nil, err)
	submittedAt, err := time.Parse(layout, "2023-03-11 06:46:57 +0000 UTC")
	assert.Equal(t, nil, err)

	assert.Equal(t, &PullRequestSummary{
		Number:   16144,
		OpenedBy: "SpaceIm", CreatedAt: createdAt, Recipe: "re2", Change: EDIT, Weight: SMALL,
		ReviewURL:     "https://github.com/conan-io/conan-center-index/pull/16144",
		LastCommitSHA: "e2aa65c961d48d688dd5450811229eb1d62649ba", LastCommitAt: lastCommitAt, CciBotPassed: false,
		Summary: Reviews{Count: 2, ValidApprovals: 2, TeamApproval: true, Approvals: []string{"toge", "prince-chrismc"},
			Blockers: nil, LastReview: &Review{ReviewerName: "prince-chrismc", SubmittedAt: submittedAt,
				HTMLURL: "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1335829632",
			},
		},
	}, review)

	assert.Equal(t, true, gock.IsDone())
}
