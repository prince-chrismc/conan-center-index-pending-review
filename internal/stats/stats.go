package stats

import (
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/duration"
)

// Stats for the open pull requests being evaluated
type Stats struct {
	Open    int
	Draft   int
	Review  int
	Merge   int
	Stopped int
	Age     duration.MovingAverage
}

// Add two stats together
func (stats *Stats) Add(s Stats) {
	stats.Age.Combine(s.Age)
	stats.Open += s.Open
	stats.Draft += s.Draft
	stats.Stopped += s.Stopped
	stats.Merge += s.Merge
	stats.Review += s.Review
}
