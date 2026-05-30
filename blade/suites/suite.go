package suites

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/auho/go-toolkit/console/output"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Suite struct {
	Cmd       *cobra.Command
	Validator *validator.Validate

	pfFunc  map[int][]func() string
	pfIndex int

	mutex sync.Mutex
}

func (se *Suite) Init(cmd *cobra.Command) {
	se.Cmd = cmd
	se.Validator = validator.New()
	se.pfFunc = make(map[int][]func() string)
}

func (se *Suite) CmdVisit(cmd *cobra.Command) string {
	var args []string
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		value := flag.Value.String()
		args = append(args, fmt.Sprintf("--%s='%s'", flag.Name, value))
	})

	return strings.Join(args, " ")
}

func (se *Suite) Watch(title string, totalDuration, interval time.Duration, fn func() ([]string, error)) {
	se.PrintlnTitle(title)

	_r, cancel := output.NewRefreshWithCancel(
		output.WithContent(fn),
		output.WithInterval(interval),
	)

	_r.Start()

	<-time.After(totalDuration)
	cancel()
}

func (se *Suite) Func(title string, fn func() error) {
	se.PrintlnTitle(title)
	err := fn()
	if err != nil {
		se.PrintlnErr(err)
	}

	se.Cmd.Println()
}
