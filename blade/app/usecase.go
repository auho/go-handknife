package app

import (
	"strings"

	"github.com/auho/go-handknife/blade/app/command/argument/parser"
	"github.com/spf13/cobra"
)

type UseCase struct{}

func (uc *UseCase) CmdFlags(cmd *cobra.Command, parsers ...parser.Parser) {
	for _, pr := range parsers {
		pr.Flags(cmd)
	}
}

func (uc *UseCase) CmdParse(parsers ...func() error) error {
	var err error
	for _, pr := range parsers {
		err = pr()
		if err != nil {
			return err
		}
	}

	return err
}

func (uc *UseCase) CmdArgsToString(s string, parsers ...parser.Parser) string {
	var ss []string
	for _, pr := range parsers {
		ss = append(ss, pr.ArgsToString()...)
	}

	if s != "" {
		ss = append(ss, s)
	}

	return strings.Join(ss, " ")
}
