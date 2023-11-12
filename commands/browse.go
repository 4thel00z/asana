package commands

import (
	"os/exec"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/4thel00z/asana/api"
	"github.com/4thel00z/asana/config"
	"github.com/4thel00z/asana/utils"
)

func Browse(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
