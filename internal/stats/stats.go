package stats

import (
	"time"
)

// Stats for the open pull requests being evaluated
type Stats struct {
	Open    int
	Draft   int
	Review  int
	Merge   int
	Stale   int
	Failed  int
	Blocked int
	Age     time.Duration
}

// Add two stats together
func (stats *Stats) Add(s Stats) {
	stats.Age = time.Duration(((stats.Age.Nanoseconds() * int64(stats.Open)) + (s.Age.Nanoseconds() * int64(s.Open))) / int64(stats.Open+s.Open))
	stats.Open += s.Open
	stats.Draft += s.Draft
	stats.Stale += s.Stale
	stats.Failed += s.Failed
	stats.Blocked += s.Blocked
	stats.Merge += s.Merge
	stats.Review += s.Review
}
