package main

import (
	"example/app"
	"example/modules"
	"github.com/auho/go-handknife/emergencybox/cmd/execute"
	"github.com/spf13/cobra"
)

var version = ""
var lastDate = ""

func main() {
	execute.Exec(func(command *cobra.Command) {
		app.Initialization()

		modules.Initialization(command)
	})
}
