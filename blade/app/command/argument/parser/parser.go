package parser

import "github.com/spf13/cobra"

type Parser interface {
	Flags(*cobra.Command)
	Parse() error
	ArgsToString() []string
}
