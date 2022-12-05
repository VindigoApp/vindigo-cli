package utils

import (
	"encoding/json"
	"os"
)

type packageJson struct {
	Dependencies map[string]string `json:"dependencies"`
}

// Returns whether the current working directory
// can be considered a valid Vindigo project
func IsWithinProject() bool {
	pkgJson, err := os.Open("package.json")

	if err != nil {
		return false
	}

	defer pkgJson.Close()

	pkg := packageJson{}

	if err := json.NewDecoder(pkgJson).Decode(&pkg); err != nil {
		return false
	}

	return pkg.Dependencies["@vindigo/server"] != ""
}
