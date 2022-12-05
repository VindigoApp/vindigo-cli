package cmd

import (
	"fmt"

	"github.com/VindigoApp/vindigo-cli/utils"
	"github.com/urfave/cli/v2"
)

func HandleStart(c *cli.Context) error {
	if !utils.IsWithinProject() {
		fmt.Println(utils.Error + "You are not in a Vindigo project")
		return nil
	}

	return nil
}
