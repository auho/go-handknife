package suites

import (
	"errors"
	"fmt"
	"testing"

	"github.com/auho/go-handknife/emergencybox/suites/template"
	"github.com/spf13/cobra"
)

func TestSuites(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test suites",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	s := &Suite{}
	s.Init(cmd)

	s.printSliceMapStringAny("test", func() ([]map[string]any, *Post, error) {
		var sm []map[string]any

		for i := 0; i < 10; i++ {
			_t := make(map[string]any)
			for j := 0; j < 5; j++ {
				_t[fmt.Sprintf("j-%d", j)] = j
			}
			sm = append(sm, _t)
		}

		return sm,
			NewPost().AddTemplate(func(t *template.Template) {
				t.AddField("j-0", func(f *template.Field) {
					f.AddPipeline(`printf "%10v"`)
					f.AddValidator(func(x any, ps template.Pipelines) (any, template.Pipelines, error) {
						return x, nil, errors.New("test error")
					})
				})
			}),
			nil
	})
}
