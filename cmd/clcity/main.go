package main

import (
	"log"
	"os"
	"time"

	"bitbucket.org/friasdesign/pfetcher/internal/commands"
	"github.com/urfave/cli"
)

const clcityDesc = `This command contains all tools needed to fetch data from 'Cuando Llega City Bus' API, generate GTFS feed from it and/or KML file to see that data on a map. This CLI app uses publicly available data from 'Cuando Llega City Bus' API without violating any security measurement, so it's 100% legal.`

func main() {
	app := cli.NewApp()
	app.Name = "clcity"
	app.Usage = "A data fetcher and parser for 'Cuando Llega City Bus' API"
	app.Description = clcityDesc
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Carlos Frias",
			Email: "carlos.a.frias@gmail.com",
		},
	}
	app.HelpName = "clcity"
	app.Commands = commands.Commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
