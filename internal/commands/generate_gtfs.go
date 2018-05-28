package commands

import (
	"fmt"
	"path"

	"github.com/friasdesign/clcity/internal/clcitybusapi/gtfs"

	"github.com/friasdesign/clcity/pkg/elapsed"
	"github.com/urfave/cli"
)

func gtfsAction(c *cli.Context) error {
	defer elapsed.Elapsed()()
	dumpPath := "testdata"
	gtfsPath := path.Join(dumpPath, "")

	emp, err := fetch(dumpPath)
	if err != nil {
		return err
	}

	fmt.Println("Generating GTFS file ...")
	err = gtfs.Generate(emp, gtfsPath)
	if err != nil {
		return err
	}
	fmt.Println("Generating GTFS file ... DONE!")

	return nil
}
