package validate

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/stats"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// OnlyAcceptableLabels checks is the labels associated to the pull request are exclusively `Docs` or `GitHub Config`
func OnlyAcceptableLabels(labels []*pending_review.Label, stats *stats.Stats) bool {
	isDoc := false
	isConfig := false
	len := len(labels)
	if len > 0 {
		for _, label := range labels {
			switch label.GetName() {
			case "Docs":
				isDoc = true
			case "GitHub config":
				isConfig = true
			case "stale":
				stats.Stale++
			case "Failed": // , "Unexpected Error" Are alway tagged failed
				stats.Failed++
			case "infrastructure", "blocked":
				stats.Blocked++
			}
		}

		if !isDoc && !isConfig {
			return false // We know if there are certain labels then there's probably something wrong!
		}
	}

	if len > 1 && !isDoc { // We always want to review documentation changes
		return false // We know if there are certain labels then it's probably something worth skipping!
	}

	return true
}
