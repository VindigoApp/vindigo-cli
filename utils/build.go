package utils

var (
	Production string
	Version    string
	Commit     string
	Date       string
)

func IsProduction() bool {
	return Production == "true"
}
