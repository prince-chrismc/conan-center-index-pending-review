package duration

import (
	"fmt"
	"strings"
	"time"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

// Duration Convert the default time.Duration to a slightly more human readable format
func Duration(d time.Duration) string {
	/// https://gist.github.com/harshavardhana/327e0577c4fed9211f65#gistcomment-2557682
	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd%s", days, d)

	return b.String()
}
