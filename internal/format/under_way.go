package format

import (
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

func UnderReview(prs []*pending_review.PullRequestSummary) string {
	tableBody, rowCount := ReviewsToMarkdownRows(prs, false)

	if rowCount == 0 {
		return `
		:confused: There's nothing within the review process... You should [open a bug report](https://github.com/prince-chrismc/conan-center-index-pending-review/issues/new)
		`
	}

	breif := "is **1** pull request"
	if rowCount > 1 {
		breif = "are ***" + fmt.Sprint(rowCount) + "** pull requests"
	}

	return `

	### :nerd_face: Please Review! 
	
	There ` + breif + ` currently under way :eyes:
	
	PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
	:---: | --- | --- | :---: | --- | ---
	` + tableBody
}
