package execute

import (
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var _ io.Writer = (*newErr)(nil)

type newErr struct{}

func (e *newErr) Write(p []byte) (n int, err error) {
	return os.Stderr.Write([]byte(text.FgHiRed.Sprint(string(p))))
}
