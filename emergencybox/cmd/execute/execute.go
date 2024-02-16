package execute

import (
	"os"
	"path"
	"runtime"

	"github.com/auho/go-handknife/emergencybox/cmd/module"
	"github.com/auho/go-handknife/emergencybox/cmd/widget"
	"github.com/spf13/cobra"
)

func Exec(fns ...func(command *cobra.Command)) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	frame, _ := runtime.CallersFrames(pc).Next()

	appName := path.Base(path.Dir(frame.File))
	if appName == "" {
		panic("appName is empty")
	}

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use: appName,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := widget.NewTime().ServerTimeContrast()
			if err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.SetErr(&newErr{})

	ExecWithRoot(appName, rootCmd, fns...)
}

func ExecWithRoot(appName string, rootCmd *cobra.Command, fns ...func(command *cobra.Command)) {
	// module command
	module.Initial(rootCmd, appName)

	for _, fn := range fns {
		fn(rootCmd)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
