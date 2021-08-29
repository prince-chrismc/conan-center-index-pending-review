package main

import (
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-open-versus-merged",
		Usage: "create a graph displaying the trend of open, merged and closed pull requests",
		Flags: app.DefaultFlags(),
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")
			token := c.String("access-token")

			return OpenVersusMerged(token, dryRun)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
