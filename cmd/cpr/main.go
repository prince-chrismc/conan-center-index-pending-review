package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-pending-review",
		Usage: "create a comprehensive list of all the open pull requests under review and how far along they are",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"d"},
				Usage:   "scrap the GitHub API for all the relevant information but do NOT post the results",
			},
			&cli.StringFlag{
				Name:    "access-token",
				Aliases: []string{"t"},
				Usage:   "a GitHub `access-token` to use, this can be either the default or a Personal Access Token (PAT).",
				EnvVars: []string{"ACCESS_TOKEN", "GITHUB_TOKEN"},
			},
		},
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")
			token := c.String("access-token")

			return PendingReview(token, dryRun)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
