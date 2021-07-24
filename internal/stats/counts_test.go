package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountAtTimeAddsValues(t *testing.T) {
	counts := CountAtTime{}
	assert.Equal(t, CountAtTime{}, counts)

	t1 := time.Now()
	counts.Add(t1, 5)
	assert.Equal(t, CountAtTime{t1: 5}, counts)

	t2 := time.Now()
	counts.Add(t2, 27)
	assert.Equal(t, CountAtTime{t1: 5, t2: 27}, counts)
}

func TestCountAtTimeIncrementsValues(t *testing.T) {
	counts := CountAtTime{}
	assert.Equal(t, CountAtTime{}, counts)

	t1 := time.Now()
	counts.Count(t1)
	assert.Equal(t, CountAtTime{t1: 1}, counts)

	counts.Count(t1)
	assert.Equal(t, CountAtTime{t1: 2}, counts)

	t2 := time.Now()
	counts.Count(t2)
	assert.Equal(t, CountAtTime{t1: 2, t2: 1}, counts)
}

func TestCountAtTimeGivesKeys(t *testing.T) {
	counts := CountAtTime{}
	t1 := time.Now()
	counts.Add(t1, 5)
	t2 := time.Now()
	counts.Count(t2)

	keys := counts.Keys()
	assert.Equal(t, []time.Time{t1, t2}, keys)
}
