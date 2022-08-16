package main

import (
	"fmt"
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-open-versus-merged",
		Usage: "create a graph displaying the trend of open, merged and closed pull requests",
		Flags: app.DefaultFlags(),
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")

			return OpenVersusMerged(dryRun, c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s failed: %v\n", app.Name, err)
		os.Exit(1)
	}
}
