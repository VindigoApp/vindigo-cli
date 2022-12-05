package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/VindigoApp/vindigo-cli/utils"
	"github.com/jwalton/gchalk"
	"github.com/urfave/cli/v2"
)

var commandMap = map[string]string{
	"npm":  "npm install",
	"yarn": "yarn install",
	"pnpm": "pnpm install",
}

func HandleInit(c *cli.Context) error {
	fmt.Println(utils.Info + "Initializing project")

	version := c.String("version")
	pkgMngr := c.String("package-manager")
	installCmd := commandMap[pkgMngr]
	folderName := c.Args().Get(0)

	// Validate package manager
	if installCmd == "" {
		fmt.Println(utils.Error + "Invalid package manager. Valid options are: npm, yarn, pnpm")
		return nil
	}

	// Use default folder name
	if folderName == "" {
		folderName = "vindigo"
	}

	// Create folder structure
	if err := os.Mkdir(folderName, os.ModePerm); err != nil {
		fmt.Println(utils.Error + "Failed to create directory")
		return err
	}

	if err := os.Chdir(folderName); err != nil {
		fmt.Println(utils.Error + "Failed to change directory")
		return err
	}

	if err := os.Mkdir("data", os.ModePerm); err != nil {
		fmt.Println(utils.Error + "Failed to create data directory")
		return err
	}

	// Find latest version
	if version == "" {
		res, err := exec.Command("npm", "show", "@vindigo/server", "version").Output()

		if err != nil {
			fmt.Println(utils.Error + "Failed to find latest version")
			return err
		}

		version = strings.TrimSpace(string(res))
	}

	// Create package json
	pkgJson, err := utils.Resources.ReadFile("package.json")

	if err != nil {
		fmt.Println(utils.Error + "Failed to read package.json")
		return err
	}

	pkgText := string(pkgJson)
	pkgText = strings.ReplaceAll(pkgText, "%PM%", pkgMngr)
	pkgText = strings.ReplaceAll(pkgText, "%VER%", version)

	if err := os.WriteFile("package.json", []byte(pkgText), 0644); err != nil {
		fmt.Println(utils.Error + "Failed to write package.json")
		return err
	}

	fmt.Println(utils.Info + "Installing server...")

	// Install dependencies
	installArgs := strings.Split(installCmd, " ")

	if err := exec.Command(installArgs[0], installArgs[1:]...).Run(); err != nil {
		fmt.Println(utils.Error + "Failed to install dependencies")
		return err
	}

	fmt.Println(utils.Succes + "Installion complete!")
	fmt.Println(utils.Succes + "Use " + gchalk.Yellow("cd "+folderName) + " to enter the project directory")

	return nil
}
