package stats

import (
	"sort"
	"time"
)

// CountAtTime provides the number of elements at each given time
type CountAtTime map[time.Time]int

// Add insert the count at time
func (c CountAtTime) Add(time time.Time, count int) {
	c[time] = count
}

// AddNow insert the count at `time.Now()`
func (c CountAtTime) AddNow(count int) {
	c.Add(time.Now(), count)
}

// Count increments the count at time
func (c CountAtTime) Count(time time.Time) {
	currentCounter, found := c[time]
	if found {
		c[time] = currentCounter + 1
	} else {
		c[time] = 1
	}
}

func (c CountAtTime) Keys() []time.Time {
	keys := make([]time.Time, len(c)) // Allocate everything in one step for performance
	idx := 0
	for time := range c {
		keys[idx] = time // Fill each elemenet of the array
		idx++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return keys
}
