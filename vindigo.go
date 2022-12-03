package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	log.Println(strings.Join(os.Args, ", "))
}
