package duration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverageCalculation(t *testing.T) {
	ma := MovingAverage{}
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(0), "initial value should be zero")

	ma.Append(time.Duration(123))
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(123), "first fill should have the same value")

	ma.Append(time.Duration(1))
	assert.Equal(t, ma.GetCurrentAverage(), time.Duration(62), "zero sized sample should half value")
}
