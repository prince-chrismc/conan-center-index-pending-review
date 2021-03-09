package duration

import (
	"fmt"
	"strings"
	"time"
)

const (
	hour = time.Minute * 60
	day  = hour * 24
	year = 365 * day
)

// Duration represents the elapsed time between two instants as a time.Duration. It adds a few extras for working with the scale of CCI
type Duration = time.Duration

// String converts the default time.Duration to a slightly more human readable format
func String(d Duration) string {
	var b strings.Builder
	and := false

	if d >= year {
		and = true
		years := d / year
		d -= years * year
		fmt.Fprintf(&b, "%d years, ", years)
	}

	if d >= day {
		and = true
		days := d / day
		d -= days * day
		fmt.Fprintf(&b, "%d days, ", days)
	}

	if d >= hour {
		and = true
		hours := d / hour
		d -= hours * hour
		fmt.Fprintf(&b, "%d hours, ", hours)
	}

	if and {
		fmt.Fprintf(&b, "and ")
	}

	fmt.Fprintf(&b, "%.02f minutes", d.Minutes())

	return b.String()
}
