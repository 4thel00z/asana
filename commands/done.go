package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/4thel00z/asana/api"
)

func Done(c *cli.Context) {
	task := api.Update(api.FindTaskId(c.Args().First(), false), "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
