package format

import (
	"fmt"
	"strings"
	"time"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
)

// ReviewsToMarkdownRows Converts pull request review status to GitHub markdown table rows
func ReviewsToMarkdownRows(prs []*pending_review.PullRequestSummary, canMerge bool) (string, int) {
	count := 0
	var retval string
	for _, pr := range prs {
		if pr.Summary.IsApproved() != canMerge {
			continue // Skip what we do not want
		}

		count++
		if canMerge {
			retval += toMerge(pr)
		} else {
			retval += underWay(pr)
		}
	}
	return retval, count
}

func underWay(pr *pending_review.PullRequestSummary) string {
	var retval string
	title := title(pr.Change, pr.Recipe)
	if !pr.CciBotPassed {
		// The assumption here is that "no passing" means in progress since "failing" PRs are disgarded
		title = ":stopwatch: " + pr.Recipe
	}

	columns := []string{
		fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
		fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
		pr.CreatedAt.Format("Jan 2"),
		title,
		fmt.Sprint(pr.Summary.Count),
		lastReviewTime(pr),
		strings.Join(pr.Summary.Blockers, ", "),
		strings.Join(pr.Summary.Approvals, ", "),
	}

	retval += strings.Join(columns, "|")
	retval += "\n"

	return retval
}

func toMerge(pr *pending_review.PullRequestSummary) string {
	var retval string
	title := title(pr.Change, pr.Recipe)
	if !pr.CciBotPassed {
		switch pr.Change {
		case pending_review.NEW, pending_review.EDIT:
			title = ":warning: " + pr.Recipe
		}
	}

	columns := []string{
		fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
		fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
		pr.CreatedAt.Format("Jan 2"),
		title,
		fmt.Sprint(pr.Summary.Count),
		strings.Join(pr.Summary.Approvals, ", "),
	}
	retval += strings.Join(columns, "|")
	retval += "\n"

	return retval
}

func title(change pending_review.Category, recipe string) string {
	switch change {
	case pending_review.NEW:
		return ":new: " + recipe
	case pending_review.EDIT:
		return ":memo: " + recipe
	case pending_review.DOCS:
		return ":green_book: " + recipe
	case pending_review.CONFIG:
		return ":gear: " + recipe
	}

	return "???"
}

func lastReviewTime(pr *pending_review.PullRequestSummary) string {
	if pr.Summary.LastReview != nil {
		date := pr.Summary.LastReview.SubmittedAt.Format("Jan 2")

		if time.Since(pr.Summary.LastReview.SubmittedAt) >= duration.DAY*12 {
			date += " :bell:"
		}

		return date
	}

	if time.Since(pr.LastCommitAt) >= duration.DAY*3 {
		return ":eyes:"
	}

	return ""
}
