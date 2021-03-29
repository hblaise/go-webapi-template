package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Run database migrations",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "import",
				Aliases: []string{"i"},
				Usage:   "Import data",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
