package format

import (
	"fmt"
	"strings"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// ReviewsToMarkdownRows Converts pull request review status to GitHub markdown table rows
func ReviewsToMarkdownRows(prs []*pending_review.PullRequestStatus, canMerge bool) string {
	var retval string
	for _, pr := range prs {
		if pr.IsMergeable != canMerge {
			continue
		}

		title := "recipe"
		switch pr.Change {
		case pending_review.ADDED:
			title = ":new: " + pr.Recipe
			break
		case pending_review.EDIT:
			title = ":memo: " + pr.Recipe
			break
		case pending_review.BUMP:
			title = ":arrow_up: " + pr.Recipe
			break
		case pending_review.DOCS:
			title = ":green_book: " + pr.Recipe
			break
		}

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
