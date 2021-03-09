package duration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverageAppend(t *testing.T) {
	ma := MovingAverage{}
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(0), "initial value should be zero")

	ma.Append(time.Duration(123))
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(123), "first fill should have the same value")

	ma.Append(time.Duration(1))
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(62), "zero sized sample should half value")
}

func TestMovingAverageCombine(t *testing.T) {
	ma := MovingAverage{}
	ma.Append(time.Duration(123))
	ma.Append(time.Duration(1))

	ma2 := MovingAverage{}
	ma2.Append(time.Duration(123))
	ma2.Append(time.Duration(1))

	ma.Combine(ma2)

	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(62))
}
