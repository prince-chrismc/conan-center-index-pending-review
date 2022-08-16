package main

import (
	"fmt"
	"os"

	"github.com/prince-chrismc/conan-center-index-pending-review/v3/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "conan-center-index-pending-review",
		Usage: "create a comprehensive list of all the open pull requests under review and how far along they are",
		Flags: app.DefaultFlags(),
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")

			return PendingReview(dryRun, c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s failed: %v\n", app.Name, err)
		os.Exit(1)
	}
}
