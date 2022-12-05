package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
				Action: cmd.HandleStatus,
			},
			{
				Name:      "init",
				Usage:     "Create a new Vindigo instance",
				ArgsUsage: "[name]",
				Action:    cmd.HandleInit,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "package-manager",
						Aliases: []string{"pm"},
						Value:   "npm",
						Usage:   "Which package manager to use",
					},
					&cli.StringFlag{
						Name:    "version",
						Aliases: []string{"v"},
						Usage:   "The version of vindigo to install",
					},
				},
			},
		},
	}

	input := os.Args

	if !utils.IsProduction() {
		fmt.Print("Enter command: ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		trim := strings.TrimSpace(text)
		input = append(os.Args[:1], strings.Split(trim, " ")...)
	}

	err := app.RunContext(ctx, input)

	if err != nil {
		log.Fatalln("Failed to run the application", err)
	}
}
