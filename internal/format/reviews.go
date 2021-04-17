package format

import (
	"fmt"
	"strings"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
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
	if !pr.CciBotPassed && pr.Summary.IsApproved() {
		title = ":warning: " + pr.Recipe
	}

	columns := []string{
		fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
		fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
		pr.CreatedAt.Format("Jan 2"),
		title,
		fmt.Sprint(pr.Summary.Count),
		strings.Join(pr.Summary.Blockers, ", "),
		strings.Join(pr.Summary.Approvals, ", "),
	}

	if pr.Summary.LastReview != nil {
		columns = append(columns, fmt.Sprint("[", pr.Summary.LastReview.ReviewerName, "](", pr.Summary.LastReview.HTMLURL, ") at ", pr.Summary.LastReview.SubmittedAt.Format("Jan 2")))
		// columns = append(columns, pr.Summary.LastReview.SubmittedAt.Format("Jan 2"))
	}

	retval += strings.Join(columns, "|")
	retval += "\n"

	return retval
}

func toMerge(pr *pending_review.PullRequestSummary) string {
	var retval string
	title := title(pr.Change, pr.Recipe)
	if !pr.CciBotPassed && pr.Summary.IsApproved() { //TODO(prince-chrismc): Always display bad commit statuses?
		title = ":warning: " + pr.Recipe
	}

	columns := []string{
		fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
		fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
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
	case pending_review.ADDED:
		return ":new: " + recipe
	case pending_review.EDIT:
		return ":memo: " + recipe
	case pending_review.BUMP:
		return ":arrow_up: " + recipe
	case pending_review.DOCS:
		return ":green_book: " + recipe
	}

	return "???"
}
