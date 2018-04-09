package commands

import (
	"fmt"
	"os"
	"path"

	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/kml"
	"bitbucket.org/friasdesign/clcity/pkg/elapsed"
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
		},
	}
}

func kmlAction(c *cli.Context) error {
	defer elapsed.Elapsed()()
	dumpPath := "testdata"

	if _, err := os.Stat(dumpPath); os.IsNotExist(err) {
		panic("testdata/ folder doesn't exist")
	}
	cli := client.NewClient(nil, "testdata")

	// Create empresa
	fmt.Println("Creating empresa ...")
	emp := clcitybusapi.NewEmpresa(355)
	fmt.Println("Creating empresa ... DONE!")

	// Fetch lineas
	fmt.Println("Fetching lineas ...")
	lin, err := cli.LineaService().LineasPorEmpresa(emp)
	if err != nil {
		return err
	}
	emp.Lineas = lin
	fmt.Println("Fetching lineas ... DONE!")

	// Fetch paradas
	fmt.Println("Fetching paradas ...")
	par, err := cli.ParadaService().ParadasPorEmpresa(emp)
	if err != nil {
		return err
	}
	emp.Paradas = par
	fmt.Println("Fetching paradas ... DONE!")

	fmt.Println("Generating KML file ...")
	kml.Generate(emp, path.Join(dumpPath, "city_bus.kml"))
	if err != nil {
		return err
	}
	fmt.Println("Generating KML file ... DONE!")

	return nil
}
