package app

import (
	"strings"

	"github.com/spf13/cobra"
)

type UseCastor interface {
	CmdFlags(*cobra.Command)
	ParseArgument() error
	CmdArgs() []string
}

type UseCase struct{}

func (uc *UseCase) RunParseArgument(fns ...func() error) error {
	var err error
	for _, fn := range fns {
		err = fn()
		if err != nil {
			return err
		}
	}

	return err
}

func (uc *UseCase) RunCmdFlags(cmd *cobra.Command, ucs ...UseCastor) {
	for _, _uc := range ucs {
		_uc.CmdFlags(cmd)
	}
}

func (uc *UseCase) RunCmdArgs(s string, ucs ...UseCastor) string {
	var ss []string
	for _, _uc := range ucs {
		ss = append(ss, _uc.CmdArgs()...)
	}

	if s != "" {
		ss = append(ss, s)
	}

	return strings.Join(ss, " ")
}
