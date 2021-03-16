package format

import (
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

func ReadyToMerge(prs []*pending_review.PullRequestSummary) string {
	tableBody, rowCount := ReviewsToMarkdownRows(prs, true)

	if rowCount == 0 {
		return ""
	}

	breif := "**1** pull request is"
	if rowCount > 1 {
		breif = "***" + fmt.Sprint(rowCount) + "** pull requests are"
	}

	return `

### :heavy_check_mark: Ready to Merge 

Currently ` + breif + ` waiting to be merged :tada:

PR | By | Recipe | Reviews | :stop_sign: Blockers | :star2: Approvers
:---: | --- | --- | :---: | --- | ---
` + tableBody
}
