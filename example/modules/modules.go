package modules

import (
	"example/modules/devops"
	"github.com/spf13/cobra"
)

func Initialization(rootCmd *cobra.Command) {
	devops.Initialization(rootCmd)
}
