package commands

import (
	"github.com/urfave/cli"
)

const generateUsage = `Outputs data from 'Cuando Llega City Bus' API to given formats:
	- 'generate kml' -> KML file with all stops and trips drawn.
	- 'generate gtfs' -> GTFS feed for any trip planner.
`

func generate() cli.Command {
	return cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   generateUsage,
		Subcommands: []cli.Command{
			{
				Name:   "kml",
				Usage:  "Outputs data from 'Cuando Llega City Bus' API to a KML file containing all stops and trips.",
				Action: kmlAction,
			},
			{
				Name:   "gtfs",
				Usage:  "Generates some GTFS files from data fetched from 'Cuando llega City Bus' API.",
				Action: gtfsAction,
			},
		},
	}
}
