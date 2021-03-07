package duration

import "time"

// MovingAverage allows for the calculation of a [cumulative moving average](https://en.wikipedia.org/wiki/Moving_average)
type MovingAverage struct {
	current time.Duration
	size    int64
}

// Append a duration of time to the current moving average
func (m *MovingAverage) Append(t time.Duration) {
	newTotal := t + time.Duration(m.size*m.current.Nanoseconds())
	m.current = time.Duration(float64(newTotal.Nanoseconds()) / float64(m.size+1))
	m.size++
}

// Combine a second moving average into the current
func (m *MovingAverage) Combine(a MovingAverage) {
	newTotal := time.Duration(m.current.Nanoseconds()*m.size) + time.Duration(a.current.Nanoseconds()*a.size)
	m.current = time.Duration(float64(newTotal.Nanoseconds()) / float64(m.size+a.size))
	m.size += a.size
}

// GetCurrentAverage returns the current aaverage field if it's non-nil, zero value otherwise.
func (m *MovingAverage) GetCurrentAverage() time.Duration {
	if m == nil {
		return 0
	}

	return m.current
}
