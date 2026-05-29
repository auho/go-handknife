package output

import (
	"sort"

	"github.com/auho/go-handknife/blade/suites/template"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Output struct {
	table table.Writer
}

func NewOutput() *Output {
	o := &Output{}
	o.table = table.NewWriter()

	return o
}

func NewOutputDefault() *Output {
	o := NewOutput()
	o.table.SetAutoIndex(true)

	return o
}

func (o *Output) Setting(fn func(table.Writer)) {
	fn(o.table)
}

func (o *Output) Output(data []map[string]any, t *template.Template) (string, error) {
	if len(data) <= 0 {
		return "", nil
	}

	return o.render(data, t)
}

func (o *Output) render(sm []map[string]any, t *template.Template) (string, error) {
	var keys []string
	for k := range sm[0] {
		if t.IsFieldExclude(k) {
			continue
		}

		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var _header table.Row
	for _, key := range keys {
		_header = append(_header, key)
	}

	o.table.AppendHeader(_header)

	for _, m := range sm {
		if m == nil {
			o.table.AppendSeparator()
		} else {
			var _row table.Row
			for _, key := range keys {
				_row = append(_row, m[key])
			}

			o.table.AppendRow(_row)
		}
	}

	return o.table.Render(), nil
}
