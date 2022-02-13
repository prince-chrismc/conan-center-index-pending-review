package validate

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// OnlyAcceptableLabels checks is the labels associated to the pull request are exclusively `Docs` or `Bump version`
func OnlyAcceptableLabels(labels []*pending_review.Label, stats *stats.Stats) bool {
	isBump := false
	isDoc := false
	len := len(labels)
	if len > 0 {
		for _, label := range labels {
			switch label.GetName() {
			case "Bump version", "Bump dependencies":
				isBump = true
			case "Docs":
				isDoc = true
			case "stale":
				stats.Stale++
			case "Failed": // , "Unexpected Error" Are alway tagged failed
				stats.Failed++
			case "infrastructure", "blocked":
				stats.Blocked++
			}
		}

		if !isBump && !isDoc {
			return false // We know if there are certain labels then there's probably something wrong!
		}
	}

	if len > 1 && !isDoc { // We always want to review documentation changes
		return false // We know if there are certain labels then it's probably something worth skipping!
	}

	return true
}
