package commands

import (
	"github.com/urfave/cli"
)

// Commands returns all commands for clcity cli app.
func Commands() []cli.Command {
	return []cli.Command{
		generate(),
	}
}
