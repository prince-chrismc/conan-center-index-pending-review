package format

import (
	"fmt"
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/stats"
)

// Statistics makes a markdown section and unordered list
func Statistics(stats stats.Stats) string {
	return `

#### :clipboard: Statistics

> :warning: These are just rough metrics counting the labels and may not reflect the actual state of pull requests

- Commit: ` + os.Getenv("GITHUB_SHA") + `
- Pull Requests:
	- Open: ` + fmt.Sprint(stats.Open) + `
	- Draft: ` + fmt.Sprint(stats.Draft) + `
	- Average Age: ` + duration.String(stats.Age.GetCurrentAverage()) + `
	- Stop Label: ` + fmt.Sprint(stats.Stopped) + `
	`
}
