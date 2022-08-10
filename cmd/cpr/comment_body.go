package main

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/format"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
)

// MakeCommentBody to render human friendly version that can be published to GitHub Issues
func MakeCommentBody(summaries []*pending_review.PullRequestSummary, stats stats.Stats, owner string, repo string) string {
	return `## :sparkles: Summary of Pull Requests Pending Review!

	### :ballot_box_with_check: Selection Criteria:
	
	- There has been at least one approval on the head commit
	- The last commit occurred after any reviews
	- Must not have a label indicating stopped or auto merge
	
	#### Legend
	
	:new: - Adding a recipe which does not yet exist<br>
	:memo: - Modification to an existing recipe<br>
	:green_book: - Documentation change <sup>[1]</sup><br>
	:gear: - GitHub configuration/workflow changes <sup>[1]</sup><br>
	:stopwatch: or :warning: - The commit status does **not** indicate success <sup>[2]</sup><br>
	:bell: - The last review was more than 12 days ago<br>
	:eyes: - It's been more than 3 days since the last commit and there are no reviews<br>
	<br>
	<sup>[1]</sup>: _closely_ matches the label<br>
	<sup>[2]</sup>: depending whether the PR is under way or ready to merge` +
		format.UnderReview(summaries, owner, repo) + format.ReadyToMerge(summaries) + format.Statistics(stats) + `
			
	[Raw JSON data](https://raw.githubusercontent.com/` + owner + "/" + repo + `/raw-data/pending-review.json)
	
	## :bar_chart: Open Versus Merged
	
	#### Legend
	
	:green_square: - Open pull requests<br>
	:red_square: - Closed pull requests<br>
	:purple_square: - Merged pull requests <sup>[1]</sup><br>
	
	![ovm](https://github.com/` + owner + "/" + repo + `/blob/raw-data/open-versus-merged.gif?raw=true)
	
	<sup>[1]</sup>: the darker bottom section indicated merged within 7 days of being opened
	
	## :hourglass: Time Spent in Review
	
	![tir](https://github.com/` + owner + "/" + repo + `/blob/raw-data/time-in-review.png?raw=true)
	
	Found this useful? Give it a :star: :pray:
	`
}
