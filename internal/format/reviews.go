package format

import (
	"fmt"
	"strings"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// ReviewsToMarkdownRows Converts pull request review status to GitHub markdown table rows
func ReviewsToMarkdownRows(prs []*pending_review.ReviewSummary, canMerge bool) string {
	var retval string
	for _, pr := range prs {
		if pr.IsMergeable != canMerge {
			continue
		}

		title := title(pr.Change, pr.Recipe)
		if !pr.CciBotPassed && pr.IsMergeable {
			title = ":warning: " + pr.Recipe
		}

		columns := []string{
			fmt.Sprint("[#", pr.Number, "](", pr.ReviewURL, ")"),
			fmt.Sprint("[", pr.OpenedBy, "](https://github.com/", pr.OpenedBy, ")"),
			title,
			fmt.Sprint(pr.Reviews),
			strings.Join(pr.HeadCommitBlockers, ", "),
			strings.Join(pr.HeadCommitApprovals, ", "),
		}
		retval += strings.Join(columns, "|")
		retval += "\n"
	}
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
