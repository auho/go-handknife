package execute

import (
	"os"

	"github.com/auho/go-handknife/blade/cmd/module"
	"github.com/auho/go-handknife/blade/cmd/widget"
	"github.com/spf13/cobra"
)

func Exec(appName string, fns ...func(command *cobra.Command)) {
	if appName == "" {
		panic("appName is empty")
	}

	var rootCmd = &cobra.Command{
		Use: appName,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := widget.NewTime().ServerTimeContrast(0)
			if err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.SetErr(&stdErr{})

	ExecWithRoot(appName, rootCmd, fns...)
}

func ExecWithRoot(appName string, rootCmd *cobra.Command, fns ...func(command *cobra.Command)) {
	module.Initial(rootCmd, appName)

	for _, fn := range fns {
		fn(rootCmd)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
