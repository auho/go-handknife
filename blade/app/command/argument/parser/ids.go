package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/auho/go-handknife/blade/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ Parser = (*Ids)(nil)

type Ids struct {
	IdsArg string
	Ids    []int

	title   string
	argName string
}

func (s *Ids) WithSetting(title, argName string) *Ids {
	s.title = title
	s.argName = argName

	return s
}

func (s *Ids) WithTitle(title string) *Ids {
	s.title = title

	return s
}

func (s *Ids) Flags(cmd *cobra.Command) {
	s.check()

	cmd.Flags().StringVar(&s.IdsArg, s.argName, "", s.title)
}

func (s *Ids) Parse() error {
	var err error

	s.check()

	s.IdsArg, err = prompt.NewTextWithSize(s.title, 16192)
	if err != nil {
		return err
	}

	err = s.handleIds()
	if err != nil {
		return err
	}

	return nil
}

func (s *Ids) ArgsToString() []string {
	s.check()

	if len(s.Ids) <= 0 {
		return nil
	}

	return []string{
		fmt.Sprintf("--%s %s", s.argName, s.IdsArg),
	}
}

func (s *Ids) InjectIds(ids []int) {
	var idsArg []string
	for _, _id := range ids {
		idsArg = append(idsArg, fmt.Sprintf("%d", _id))
	}

	s.Ids = ids
	s.IdsArg = strings.Join(idsArg, ",")
}

func (s *Ids) InjectArg(argIds string) error {
	s.IdsArg = argIds

	return s.handleIds()
}

func (s *Ids) handleIds() error {
	var err error

	s.IdsArg = strings.ReplaceAll(s.IdsArg, "\n", ",")
	s.IdsArg = strings.ReplaceAll(s.IdsArg, "/", ",")
	s.IdsArg = strings.ReplaceAll(s.IdsArg, "-", ",")
	s.IdsArg = strings.ReplaceAll(s.IdsArg, "_", ",")
	s.IdsArg = strings.ReplaceAll(s.IdsArg, " ", ",")

	idsString := strings.Split(s.IdsArg, ",")
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
		s.Ids = append(s.Ids, id)
	}

	return nil
}

func (s *Ids) check() {
	if s.title == "" {
		s.title = "ids"
	}

	if s.argName == "" {
		s.argName = "ids"
	}
}
