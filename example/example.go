package main

import (
	"example/app"
	"example/modules"

	"github.com/auho/go-handknife/blade/cmd/execute"
	"github.com/spf13/cobra"
)

var appName string
var version = ""
var lastDate = ""

func main() {
	appName = "example"

	execute.Exec(appName, func(command *cobra.Command) {
		app.Initialize()

		modules.Initialize(command)
	})
}
