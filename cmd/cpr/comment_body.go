package main

import (
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/format"
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
)

// MakeCommentBody to render human friendly version that can be published to GitHub Issues
func MakeCommentBody(summaries []*pending_review.PullRequestSummary, stats stats.Stats, totals stats.CountAtTime, owner string, repo string) string {
	return `## :sparkles: Summary of Pull Requests Pending Review!

### :ballot_box_with_check: Selection Criteria:

- There has been at least one approval on the head commit
- The last commit occurred after any reviews
- Must not have a label indicating stopped or auto merge

#### Legend

Icon | Description
:---:|:---
:new: | Adding a recipe which does not yet exist 
:memo: | Modification to an existing recipe 
:green_book: | Documentation change <sup>[1]</sup> 
:gear: | GitHub configuration/workflow changes <sup>[1]</sup>
:stopwatch: or :warning: | The commit status does **not** indicate success <sup>[2]</sup> 
:bell: | The last review was more than 12 days ago 
:eyes: | It's been more than 3 days since the last commit and there are no reviews 

- <sup>[1]</sup>: _closely_ matches the label
- <sup>[2]</sup>: depending whether the PR is under way or ready to merge` +
		format.UnderReview(summaries, owner, repo) + format.ReadyToMerge(summaries) + format.Statistics(stats) + `

[Raw JSON data](https://raw.githubusercontent.com/` + owner + "/" + repo + `/raw-data/pending-review.json)

## :bar_chart: Open Versus Merged

#### Legend

Pull requests are depicted as:

- Open  :green_square:
- Closed :red_square:
- Merged :purple_square:
  - Darker bottom section indicated merged within 7 days of being opened

For reference:

- 100% is ` + fmt.Sprint(totals.Values()[0]) + ` (most in the last year)
- 60% is ` + fmt.Sprint(int(float64(totals.Values()[0])*0.6)) + `

![ovm](https://github.com/` + owner + "/" + repo + `/blob/raw-data/open-versus-merged.gif?raw=true)

## :hourglass: Time Spent in Review

If you are wondering how long it will take for you pull requrest to get merged; this graph should give you an idea.

![tir](https://github.com/` + owner + "/" + repo + `/blob/raw-data/time-in-review.png?raw=true)

Found this useful? Give it a :star: :pray:
	`
}
