package utils

import (
	"embed"
)

//go:embed resources
var Resources embed.FS

var (
	Production string
	Version    string
	Commit     string
	Date       string
)

func IsProduction() bool {
	return Production == "true"
}
