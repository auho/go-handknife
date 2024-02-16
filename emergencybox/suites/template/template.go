package template

import (
	"bytes"
	"log"
	"text/template"

	"github.com/auho/go-handknife/emergencybox/suites/verbose"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Template struct {
	fields       map[string]*Field
	fieldVerbose *verbose.KeysVerboseExecute

	outputSettings []func(table.Writer)
	excludeFields  map[string]struct{}
}

func NewTemplate() *Template {
	t := &Template{}
	t.fields = make(map[string]*Field)
	t.excludeFields = make(map[string]struct{})

	return t
}

func (t *Template) IsExcludeField(field string) bool {
	_, ok := t.excludeFields[field]

	return ok
}

func (t *Template) AddExcludeField(fields ...string) *Template {
	for _, field := range fields {
		t.excludeFields[field] = struct{}{}
	}

	return t
}

func (t *Template) NewField(name string) *Field {
	if _, ok := t.fields[name]; !ok {
		t.fields[name] = &Field{}
	}

	return t.fields[name]
}

func (t *Template) AddField(name string, fn func(f *Field)) *Template {
	fn(t.NewField(name))

	return t
}

func (t *Template) AddFieldVerbose(v map[string]verbose.Verbose) *Template {
	t.fieldVerbose = verbose.NewKeysVerboseExecute(v)

	return t
}

func (t *Template) AddOutputSetting(setting func(table.Writer)) *Template {
	t.outputSettings = append(t.outputSettings, setting)

	return t
}

func (t *Template) OutputSettings() []func(writer table.Writer) {
	return t.outputSettings
}

func (t *Template) FieldExec(key string, value any) any {
	field, ok := t.fields[key]
	if ok {
		var _err string
		_value, ps, err := field.Validate(value)
		if err != nil {
			_err = " " + text.FgHiRed.Sprint(err.Error())
		}

		var _pipeline string
		if ps == nil {
			_pipeline = ""
		} else {
			_pipeline = ps.Render()
		}

		if _pipeline == "" {
			_pipeline = "{{.}}"
		} else {
			_pipeline = "{{. | " + _pipeline + "}}"
		}

		tmpl, err := template.New("").Parse(_pipeline + _err)
		if err != nil {
			log.Fatal(err, tmpl)
		}

		var b bytes.Buffer
		err = tmpl.Execute(&b, _value)
		if err != nil {
			log.Fatal(err, tmpl)
		}

		return b.String()
	} else {
		if t.fieldVerbose != nil {
			return t.fieldVerbose.ExecuteKv(key, value)
		}

		return value
	}
}
