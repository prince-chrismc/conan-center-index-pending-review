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
		// The assumption here is that "no passing" means in progress since "failing" PRs are discarded
		title = ":stopwatch: " + pr.Recipe
	}

	columns := []string{
		fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
		fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
		pr.CreatedAt.Format("Jan 2"),
		title,
		weight(pr.Weight),
		fmt.Sprint(pr.Summary.Count),
		lastReviewTime(pr),
		strings.Join(pr.Summary.Blockers, ", "),
		approvers(pr.Summary.Approvals),
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
		approvers(pr.Summary.Approvals),
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

func weight(w pending_review.ReviewWeight) string {
	switch w {
	case pending_review.TINY:
		return ":green_circle: XS"
	case pending_review.SMALL:
		return ":blue_square: S"
	case pending_review.REGULAR:
		return "M"
	case pending_review.HEAVY:
		return "L"
	case pending_review.TOO_MUCH:
		return "XL"
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

func approvers(approvers []pending_review.Approver) string {
	var names []string
	for _, a := range approvers {
		if a.Tier == pending_review.Team {
			names = append(names, fmt.Sprint(`<span style="color: #3fb950;">`, a.Name, `</span>`))

		} else {
			names = append(names, a.Name)
		}
	}

	return strings.Join(names, ", ")
}
