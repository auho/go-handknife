package app

import "github.com/spf13/cobra"

var fns []func(*cobra.Command)

func RegisterCmdInit(fn func(*cobra.Command)) {
	fns = append(fns, fn)
}

func InitParentCmd(parent *cobra.Command) {
	for _, fn := range fns {
		fn(parent)
	}
}
