package execute

import (
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var _ io.Writer = (*stdErr)(nil)

type stdErr struct{}

func (e *stdErr) Write(p []byte) (n int, err error) {
	return os.Stderr.Write([]byte(text.FgHiRed.Sprint(string(p))))
}
