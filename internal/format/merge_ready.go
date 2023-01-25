package format

import (
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
)

// ReadyToMerge formats the pull request summaries into a markdown table for those considered 'ready to merge'
func ReadyToMerge(prs []*pending_review.PullRequestSummary) string {
	tableBody, rowCount := ReviewsToMarkdownRows(prs, true)

	if rowCount == 0 {
		return ""
	}

	brief := "**1** pull request is"
	if rowCount > 1 {
		brief = "**" + fmt.Sprint(rowCount) + "** pull requests are"
	}

	return `

### :heavy_check_mark: Ready to Merge 

Currently ` + brief + ` waiting to be merged :tada:


PR | By | Opened | Recipe | Reviews | :star2: Approvers
:---: | --- | --- | --- | :---: | ---
` + tableBody
}
