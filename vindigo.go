package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VindigoApp/vindigo-cli/cmd"
	"github.com/VindigoApp/vindigo-cli/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.Background()

	app := &cli.App{
		Name:  "vindigo",
		Usage: "Manage your Vindigo installations",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print the version information",
				Action: func(c *cli.Context, arg bool) error {
					fmt.Printf("version: %s\n", utils.Version)
					fmt.Printf("commit:  %s\n", utils.Commit)
					fmt.Printf("date:    %s\n", utils.Date)
					os.Exit(0)
					return nil
				},
			},
		},
		// Before: func(c *cli.Context) error {
		// 	dataPath := c.Value("data-folder").(string)

		// 	if _, err := os.Stat(dataPath); err != nil {
		// 		return errors.New("Data folder not found (" + dataPath + ")")
		// 	}

		// 	return nil
		// },
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "Start the vindigo server in the background",
				Action: cmd.HandleStart,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "attach",
						Aliases: []string{"a"},
						Value:   false,
						Usage:   "Run the server in the current shell",
					},
				},
			},
			{
				Name:   "stop",
				Usage:  "Stop the running server",
				Action: cmd.HandleStop,
			},
			{
				Name:   "status",
				Usage:  "Stop the running server",
				Action: cmd.HandleStop,
			},
			{
				Name:   "init",
				Usage:  "Create a new Vindigo instance",
				Action: cmd.HandleStop,
			},
		},
	}

	err := app.RunContext(ctx, os.Args)

	if err != nil {
		log.Fatalln("Failed to run the application", err)
	}
}
