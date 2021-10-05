package duration

import (
	"fmt"
	"strings"
	"time"
)

// Units of Durations
const (
	HOUR  = time.Minute * 60
	DAY   = HOUR * 24
	WEEK  = 7 * DAY
	MONTH = 30 * DAY
	YEAR  = 365 * DAY
)

// Duration represents the elapsed time between two instants as a time.Duration. It adds a few extras for working with the scale of CCI
type Duration = time.Duration

// String converts the default time.Duration to a slightly more human readable format
func String(d Duration) string {
	var b strings.Builder
	and := false

	if d >= YEAR {
		and = true
		years := d / YEAR
		d -= years * YEAR
		fmt.Fprintf(&b, "%d years, ", years)
	}

	if d >= DAY {
		and = true
		days := d / DAY
		d -= days * DAY
		fmt.Fprintf(&b, "%d days, ", days)
	}

	if d >= HOUR {
		and = true
		hours := d / HOUR
		d -= hours * HOUR
		fmt.Fprintf(&b, "%d hours, ", hours)
	}

	if and {
		fmt.Fprintf(&b, "and ")
	}

	fmt.Fprintf(&b, "%.02f minutes", d.Minutes())

	return b.String()
}
