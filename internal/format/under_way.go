package format

import (
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// UnderReview formats the pull request summaries into a markdown table for those **not** considered 'ready to merge'
func UnderReview(prs []*pending_review.PullRequestSummary) string {
	tableBody, rowCount := ReviewsToMarkdownRows(prs, false)

	if rowCount == 0 {
		return `
		:confused: There's nothing within the review process... You should [open a bug report](https://github.com/prince-chrismc/conan-center-index-pending-review/issues/new)
		`
	}

	breif := "is **1** pull request"
	if rowCount > 1 {
		breif = "are **" + fmt.Sprint(rowCount) + "** pull requests"
	}

	return `

### :nerd_face: Please Review! 

There ` + breif + ` currently under way :detective:

PR | By | Opened | Recipe | Reviews | Last | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | --- | :---: | --- | --- | ---
` +
		tableBody
}
