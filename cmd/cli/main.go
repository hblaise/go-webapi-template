package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			//Migrate command example with sub commands.
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Run database migrations",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "Run migrations.",
						Action: func(c *cli.Context) error {
							fmt.Println("Running migrations...")
							return nil
						},
					},
					{
						Name:  "down",
						Usage: "Revert migrations.",
						Action: func(c *cli.Context) error {
							fmt.Println("Reverting migrations...")
							return nil
						},
					},
				},
			},
			//Import command example.
			{
				Name:    "import",
				Aliases: []string{"i"},
				Usage:   "Import data",
				Action: func(c *cli.Context) error {
					fmt.Println("Running import...")
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
