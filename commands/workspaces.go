package commands

import (
	"fmt"

	"github.com/4thel00z/asana/api"
	"github.com/urfave/cli/v2"
)

func Workspaces(c *cli.Context) {
	for _, w := range api.Me().Workspaces {
		fmt.Printf("%16d %s\n", w.Id, w.Name)
	}
}
