package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
	"github.com/urfave/cli/v2"
)

// PendingReview analysis of open pull requests
func PendingReview(dryRun bool, c *cli.Context) error {
	context := context.Background()
	client, err := internal.MakeClient(context, c)
	if err != nil {
		return fmt.Errorf("problem making client %w", err)
	}

	defer fmt.Println("::endgroup") // Always print when we return

	fmt.Println("::group::👤 Initializing list of known reviewers")
	reviewers, err := pending_review.DownloadKnownReviewersList(context, client)
	if err != nil {
		return fmt.Errorf("problem getting list of known reviewers from CCI %w", err)
	}
	fmt.Printf("%+v\n", reviewers)
	fmt.Println("::endgroup")

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	var stats stats.Stats
	var summaries []*pending_review.PullRequestSummary
	opt := &pending_review.PullRequestListOptions{
		Sort:      "created",
		Direction: "asc",
		ListOptions: pending_review.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}

	for {
		pulls, resp, err := client.PullRequests.List(context, "conan-io", "conan-center-index", opt)
		if err != nil {
			return fmt.Errorf("problem getting pull request list %w", err)
		}

		out, s, err := gatherReviewStatus(context, client, reviewers, pulls)
		if err != nil {
			return fmt.Errorf("problem getting review status %w", err)
		}

		summaries = append(summaries, out...)
		stats.Add(s)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	commentBody := MakeCommentBody(summaries, stats, client.WorkingRepository.Owner, client.WorkingRepository.Name)

	if dryRun {
		fmt.Println(commentBody)
		return nil
	}

	isDifferent, err := internal.UpdateJSONFile(context, client, "pending-review.json", summaries)
	if err != nil {
		return fmt.Errorf("problem updating 'pending-review.json' %w", err)
	}

	// https://github.com/prince-chrismc/conan-center-index-pending-review/issues/5#issuecomment-754112342
	if !isDifferent {
		fmt.Println("🦺 The obtained content is identical to the new result.")
		return nil // The published results are the same, no need to update the table.
	}

	err = updateCountFile(context, client, "review-count.json", len(summaries))
	if err != nil {
		return fmt.Errorf("problem updating 'review-count.json' %w", err)
	}

	err = updateCountFile(context, client, "open-count.json", stats.Open)
	if err != nil {
		return fmt.Errorf("problem updating 'open-count.json' %w", err)
	}

	_, err = internal.UpdateFileAtRef(context, client, "index.md", "gh-pages", []byte(commentBody))
	if err != nil {
		return fmt.Errorf("problem editing web view %w", err)
	}

	_, _, err = client.Issues.Edit(context, client.WorkingRepository.Owner, client.WorkingRepository.Name, 1, &github.IssueRequest{
		Body: github.String(commentBody),
	})
	if err != nil {
		return fmt.Errorf("problem editing issue %w", err)
	}

	fmt.Println("::endgroup")

	return nil
}

func updateCountFile(context context.Context, client *pending_review.Client, file string, count int) error {
	counts := stats.CountAtTime{}
	err := internal.GetJSONFile(context, client, file, &counts)
	if err != nil {
		return err
	}

	counts.AddNow(count)

	_, err = internal.UpdateJSONFile(context, client, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func gatherReviewStatus(context context.Context, client *pending_review.Client, reviewers *pending_review.ConanCenterReviewers, prs []*pending_review.PullRequest) ([]*pending_review.PullRequestSummary, stats.Stats, error) {
	var stats stats.Stats
	var out []*pending_review.PullRequestSummary
	for _, pr := range prs {
		stats.Age.Append(time.Since(pr.GetCreatedAt()))
		stats.Open++

		if pr.GetDraft() {
			stats.Draft++
			continue // Let's skip these
		}

		review, _, err := client.PullRequest.GetReviewSummary(context, "conan-io", "conan-center-index", reviewers, pr)
		if errors.Is(err, pending_review.ErrStoppedLabel) {
			stats.Stopped++
			continue
		} else if errors.Is(err, pending_review.ErrNoReviews) || errors.Is(err, pending_review.ErrInvalidChange) || errors.Is(err, pending_review.ErrBumpLabel) {
			continue
		} else if err != nil {
			return nil, stats, fmt.Errorf("problem getting list of reviews %w", err)
		}

		if review.Summary.IsApproved() {
			stats.Merge++
		} else {
			stats.Review++
		}

		fmt.Printf("%+v\n", review)
		out = append(out, review)
	}
	return out, stats, nil
}
