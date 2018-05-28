package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/friasdesign/clcity/internal/clcitybusapi/gtfs"
	"github.com/friasdesign/clcity/internal/clcitybusapi/kml"
	"github.com/friasdesign/clcity/pkg/elapsed"
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
		Action:  generateAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "next, n",
				Usage: "Whether to use next trips or current trips.",
			},
			cli.StringFlag{
				Name:  "path, p",
				Usage: "The path for output, defaults to 'testdata', make sure the path exists.",
			},
		},
		ArgsUsage: `
			kml -> Outputs a KML file with all stops and trips drawn.
			gtfs -> Outputs data to a GTFS feed for any trip planner.
		`,
	}
}

func generateAction(c *cli.Context) error {
	defer elapsed.Elapsed()()

	// Get dump path
	dumpPath := "testdata"
	if sPath := c.String("path"); sPath != "" {
		dumpPath = sPath
	}

	// Get empresa
	nEmp := 355
	if c.Bool("next") {
		nEmp = 450
	}

	emp, err := fetch(nEmp, dumpPath)
	if err != nil {
		return err
	}

	// Get argument
	arg0 := c.Args().Get(0)

	switch arg0 {
	case "kml":
		fmt.Println("Generating KML file ...")
		err = kml.Generate(emp, path.Join(dumpPath, "city_bus.kml"))
		fmt.Println("Generating KML file ... DONE!")
	case "gtfs":
		fmt.Println("Generating GTFS file ...")
		feedPath := path.Join(dumpPath, "feed")
		err = os.MkdirAll(feedPath, os.ModePerm)
		if err != nil {
			return err
		}
		err = gtfs.Generate(emp, feedPath)
		fmt.Println("Generating GTFS file ... DONE!")
	default:
		return fmt.Errorf("Unknown output type: '%s'", arg0)
	}

	if err != nil {
		return err
	}

	return nil
}
