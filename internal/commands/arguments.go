package commands

import (
	"os"
	"path"

	"github.com/urfave/cli"
)

func getDumpPath(c *cli.Context) (string, error) {
	path := "testdata"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errDumpDirDoesntExist
	}

	return path, nil
}

func getFeedPath(c *cli.Context) (string, error) {
	path := path.Join("testdata", "feed")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errFeedDirDoesntExist
	}

	return path, nil
}
