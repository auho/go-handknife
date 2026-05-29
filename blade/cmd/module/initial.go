package module

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Initial(rootCmd *cobra.Command, appName string) {
	var _module, _cmd, _subCmd string
	var _moduleCmd = &cobra.Command{
		Use: "module",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := NewSub(appName).Build(_module, _cmd, _subCmd)
			if err != nil {
				return fmt.Errorf("build module error: %s", err.Error())
			}

			return nil
		},
	}

	_moduleCmd.Flags().StringVarP(&_module, "module", "m", "", "module name")
	_moduleCmd.Flags().StringVarP(&_cmd, "cmd", "c", "", "command name")
	_moduleCmd.Flags().StringVarP(&_subCmd, "sub", "s", "", "sub command")
	_ = _moduleCmd.MarkFlagRequired("module")
	_ = _moduleCmd.MarkFlagRequired("cmd")

	rootCmd.AddCommand(_moduleCmd)
}
