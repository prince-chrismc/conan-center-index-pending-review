package app

import "github.com/urfave/cli/v2"

// DefaultFlags for conan-center-index-pending-review apps
func DefaultFlags() []cli.Flag {
	return []cli.Flag{
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
	}
}
