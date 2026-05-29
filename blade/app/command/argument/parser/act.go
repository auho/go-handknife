package parser

import (
	"errors"
	"fmt"

	"github.com/auho/go-handknife/blade/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ Parser = (*Act)(nil)

type Act struct {
	Act string

	title        string
	argName      string
	actSelection []string
}

func (a *Act) WithSetting(title, argName string, selections []string) *Act {
	a.title = title
	a.argName = argName
	a.actSelection = selections

	return a
}

func (a *Act) WithSettingSelections(selections []string) *Act {
	a.actSelection = selections

	return a
}

func (a *Act) Flags(cmd *cobra.Command) {
	a.check()

	cmd.Flags().StringVar(&a.Act, a.argName, "", a.title)
}

func (a *Act) Parse() error {
	var err error

	a.check()

	if len(a.actSelection) <= 0 {
		return errors.New("act selection list is empty")
	}

	a.Act, err = prompt.NewSelectWithValue(a.title, a.Act, a.actSelection)
	if err != nil {
		return err
	}

	return nil
}

func (a *Act) ArgsToString() []string {
	a.check()

	if a.Act == "" {
		return nil
	}

	return []string{
		fmt.Sprintf("--%s %s", a.argName, a.Act),
	}
}

func (a *Act) check() {
	if a.title == "" {
		a.title = "act"
	}

	if a.argName == "" {
		a.argName = "act"
	}
}
