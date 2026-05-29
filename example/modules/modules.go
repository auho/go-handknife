package modules

import (
	"example/modules/devops"

	"github.com/spf13/cobra"
)

func Initialize(rootCmd *cobra.Command) {
	devops.Initialize(rootCmd)
}
