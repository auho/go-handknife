package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/auho/go-handknife/emergencybox/app"
	"github.com/auho/go-handknife/emergencybox/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ app.UseCastor = (*IdsUseCase)(nil)

type IdsUseCase struct {
	IdsArg string
	Ids    []int

	title   string
	argName string
}

func (iuc *IdsUseCase) WithSetting(title, argName string) *IdsUseCase {
	iuc.title = title
	iuc.argName = argName

	return iuc
}

func (iuc *IdsUseCase) WIthTitle(title string) *IdsUseCase {
	iuc.title = title

	return iuc
}

func (iuc *IdsUseCase) ParseArgument() error {
	var err error

	iuc.check()

	iuc.IdsArg, err = prompt.NewTextWithSize(iuc.title, 16192)
	if err != nil {
		return err
	}

	err = iuc.parseIds()
	if err != nil {
		return err
	}

	return nil
}

func (iuc *IdsUseCase) CmdFlags(cmd *cobra.Command) {
	iuc.check()

	cmd.Flags().StringVar(&iuc.IdsArg, iuc.argName, "", iuc.title)
}

func (iuc *IdsUseCase) CmdArgs() []string {
	iuc.check()

	if len(iuc.Ids) <= 0 {
		return nil
	}

	return []string{
		fmt.Sprintf("--%s %s", iuc.argName, iuc.IdsArg),
	}
}

func (iuc *IdsUseCase) InjectIds(ids []int) {
	var idsArg []string
	for _, _id := range ids {
		idsArg = append(idsArg, fmt.Sprintf("%d", _id))
	}

	iuc.Ids = ids
	iuc.IdsArg = strings.Join(idsArg, ",")
}

func (iuc *IdsUseCase) InjectArg(argIds string) error {
	iuc.IdsArg = argIds

	return iuc.parseIds()
}

func (iuc *IdsUseCase) parseIds() error {
	var err error

	iuc.IdsArg = strings.ReplaceAll(iuc.IdsArg, "\n", ",")
	iuc.IdsArg = strings.ReplaceAll(iuc.IdsArg, "/", ",")
	iuc.IdsArg = strings.ReplaceAll(iuc.IdsArg, "-", ",")
	iuc.IdsArg = strings.ReplaceAll(iuc.IdsArg, "_", ",")
	iuc.IdsArg = strings.ReplaceAll(iuc.IdsArg, " ", ",")

	idsString := strings.Split(iuc.IdsArg, ",")
	idsFlag := make(map[int]struct{})
	for _, idString := range idsString {
		var id int
		id, err = strconv.Atoi(strings.TrimSpace(idString))
		if err != nil {
			return err
		}

		if id <= 0 {
			continue
		}

		if _, ok := idsFlag[id]; ok {
			continue
		}

		idsFlag[id] = struct{}{}
		iuc.Ids = append(iuc.Ids, id)
	}

	return nil
}

func (iuc *IdsUseCase) check() {
	if iuc.title == "" {
		iuc.title = "ids"
	}

	if iuc.argName == "" {
		iuc.argName = "ids"
	}
}
