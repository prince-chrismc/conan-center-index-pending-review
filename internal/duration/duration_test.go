package duration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationString(t *testing.T) {
	time := time.Minute + time.Second*30

	assert.Equal(t, String(time), "1.50 minutes")

	time += HOUR * 5
	assert.Equal(t, String(time), "5 hours, and 1.50 minutes")

	time += DAY * 27
	assert.Equal(t, String(time), "27 days, 5 hours, and 1.50 minutes")

	time += YEAR * 2
	assert.Equal(t, String(time), "2 years, 27 days, 5 hours, and 1.50 minutes")
}
