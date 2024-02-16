package usecase

import (
	"errors"
	"fmt"

	"github.com/auho/go-handknife/emergencybox/app"
	"github.com/auho/go-handknife/emergencybox/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ app.UseCastor = (*ActUseCase)(nil)

type ActUseCase struct {
	Act string

	title        string
	argName      string
	actSelection []string
}

func (auc *ActUseCase) WithSetting(title, argName string, selections []string) *ActUseCase {
	auc.title = title
	auc.argName = argName
	auc.actSelection = selections

	return auc
}

func (auc *ActUseCase) WithSettingSelections(selections []string) *ActUseCase {
	auc.actSelection = selections

	return auc
}

func (auc *ActUseCase) ParseArgument() error {
	var err error

	auc.check()

	if len(auc.actSelection) <= 0 {
		return errors.New("act selection list is empty")
	}

	auc.Act, err = prompt.NewSelectWithValue(auc.title, auc.Act, auc.actSelection)
	if err != nil {
		return err
	}

	return nil
}

func (auc *ActUseCase) CmdArgs() []string {
	auc.check()

	if auc.Act == "" {
		return nil
	}

	return []string{
		fmt.Sprintf("--%s %s", auc.argName, auc.Act),
	}
}

func (auc *ActUseCase) CmdFlags(cmd *cobra.Command) {
	auc.check()

	cmd.Flags().StringVar(&auc.Act, auc.argName, "", auc.title)
}

func (auc *ActUseCase) check() {
	if auc.title == "" {
		auc.title = "act"
	}

	if auc.argName == "" {
		auc.argName = "act"
	}
}
