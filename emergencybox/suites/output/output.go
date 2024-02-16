package output

import (
	"sort"

	"github.com/auho/go-handknife/emergencybox/suites/template"
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

func (o *Output) setSetting(setting func(table.Writer)) {
	o.table.SetAutoIndex(true)

	setting(o.table)
}

func (o *Output) Output(data []map[string]any, t *template.Template, setting func(writer table.Writer)) (string, error) {
	if len(data) <= 0 {
		return "", nil
	}

	o.setSetting(setting)

	return o.Render(data, t)
}

func (o *Output) Render(sm []map[string]any, t *template.Template) (string, error) {
	var keys []string
	for k := range sm[0] {
		if t.IsExcludeField(k) {
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
