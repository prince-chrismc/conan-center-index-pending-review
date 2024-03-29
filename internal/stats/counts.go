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

// Keys returns all the time values recorded sorted
func (c CountAtTime) Keys() []time.Time {
	keys := make([]time.Time, 0, len(c))
	for time := range c {
		keys = append(keys, time) // Fill each elemenet of the array
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return keys
}

// Values returns all the recorded counters sorted
func (c CountAtTime) Values() []int {
	values := make([]int, 0, len(c))
	for _, count := range c {
		values = append(values, count) // Fill each elemenet of the array
	}

	sort.SliceStable(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values
}
