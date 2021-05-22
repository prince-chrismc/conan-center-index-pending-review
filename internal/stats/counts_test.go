package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationString(t *testing.T) {
	counts := CountAtTime{}
	assert.Equal(t, counts, CountAtTime{})

	t1 := time.Now()
	counts.Add(t1, 5)
	assert.Equal(t, counts, CountAtTime{t1: 5})

	t2 := time.Now()
	counts.Add(t2, 27)
	assert.Equal(t, counts, CountAtTime{t1: 5, t2: 27})
}
