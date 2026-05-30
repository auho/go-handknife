package suites

import (
	"errors"
	"fmt"
	"testing"

	"github.com/auho/go-handknife/blade/suites/template"
	"github.com/spf13/cobra"
)

func TestPfSliceToKV(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test pf slice to kv",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	s := &Suite{}
	s.Init(cmd)

	s.PfSliceToKV("test odd length", func() ([]any, *Post, error) {
		return []any{"key1", "val1", "key2"}, NewPost(), nil
	})

	s.PfSliceToKV("test even length", func() ([]any, *Post, error) {
		return []any{"key1", "val1", "key2", "val2"}, NewPost(), nil
	})

	s.PfSliceToKV("test empty", func() ([]any, *Post, error) {
		return []any{}, NewPost(), nil
	})

	s.PfSliceToKV("test single element", func() ([]any, *Post, error) {
		return []any{"key1"}, NewPost(), nil
	})
}

func TestSuites(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test suites",
		Run: func(cmd *cobra.Command, args []string) {},
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
			NewPost().SettingTemplate(func(t *template.Template) {
				t.AddField("j-0", func(f *template.Field) {
					f.AddPipeline(`printf "%10v"`)
					f.AddDetailed(func(x any) (any, error) {
						return x, errors.New("test error")
					})
				})
			}),
			nil
	})
}
