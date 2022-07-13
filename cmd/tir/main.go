package main

import (
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-time-in-review",
		Usage: "create a comprehensive list of all the open pull requests under review and how far along they are",
		Flags: app.DefaultFlags(),
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")
			token := c.String("access-token")
			owner := c.String("repo-owner")
			repo := c.String("repo-name")

			return TimeInReview(token, dryRun, owner, repo)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
