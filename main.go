package main

import (
	"github.com/joshmalbrecht/note/cmd"
	"github.com/joshmalbrecht/note/internal/config"
)

var (
	version = "devel"
	commit  = "none"
	date    = "unknown"
)

func main() {
	config.Version = version
	config.Commit = commit
	config.Date = date

	cmd.Execute()
}
