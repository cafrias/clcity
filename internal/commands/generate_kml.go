package commands

import (
	"fmt"
	"path"

	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/kml"
	"bitbucket.org/friasdesign/clcity/pkg/elapsed"
	"github.com/urfave/cli"
)

func kmlAction(c *cli.Context) error {
	defer elapsed.Elapsed()()

	dumpPath, err := getDumpPath(c)
	if err != nil {
		return err
	}

	emp, err := fetch(dumpPath)
	if err != nil {
		return err
	}

	fmt.Println("Generating KML file ...")
	err = kml.Generate(emp, path.Join(dumpPath, "city_bus.kml"))
	if err != nil {
		return err
	}
	fmt.Println("Generating KML file ... DONE!")

	return nil
}