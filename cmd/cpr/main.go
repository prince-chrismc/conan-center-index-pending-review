package main

import (
	"fmt"
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-pending-review",
		Usage: "create a comprehensive list of all the open pull requests under review and how far along they are",
		Flags: append(app.DefaultFlags(), &cli.UintFlag{
			Name:    "single-pr",
			Aliases: []string{"pr"},
			Usage:   "Optional value of a single pull request to run the analysis over",
		},
		&cli.IntFlag{
			Name:    "issue-number",
			Aliases: []string{"i"},
			Value:   0,
			Usage:   "The number of the issue to update with the summary.",
		}),
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")
			token := c.String("access-token")
			owner := c.String("repo-owner")
			repo := c.String("repo-name")
			issue := c.Int("issue-number")

			pr := c.Uint("single-pr")
			if pr != 0 {
				return SingleReviewStatus(token, pr, owner, repo)
			}

			return PendingReview(token, dryRun, owner, repo, issue)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s failed: %v\n", app.Name, err)
		os.Exit(1)
	}
}
