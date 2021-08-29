package stats

import "time"

// DurationAtTime provides the of period incurred at each given time
type DurationAtTime map[time.Time]time.Duration
