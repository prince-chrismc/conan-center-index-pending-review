package stats

import "time"

// CountAtTime provides the number of elements at each given time
type CountAtTime map[time.Time]int

func (c CountAtTime) Add(t time.Time, count int) {
	c[t] = count
}

func (c CountAtTime) AddNow(count int) {
	c.Add(time.Now(), count)
}
