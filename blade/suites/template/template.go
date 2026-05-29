package template

import (
	"github.com/auho/go-handknife/blade/suites/verbose"
)

type Template struct {
	fields        map[string]*Field
	fieldsVerbose *verbose.Keys

	fieldsExclude map[string]struct{}
}

func NewTemplate() *Template {
	t := &Template{
		fields:        make(map[string]*Field),
		fieldsVerbose: verbose.NewKeys(),
		fieldsExclude: make(map[string]struct{}),
	}

	return t
}

func (t *Template) IsFieldExclude(field string) bool {
	_, ok := t.fieldsExclude[field]

	return ok
}

func (t *Template) AddFieldsExclude(fields ...string) *Template {
	for _, field := range fields {
		t.fieldsExclude[field] = struct{}{}
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

func (t *Template) AddFieldVerbose(key string, v verbose.Verbose) *Template {
	t.fieldsVerbose.Add(key, v)

	return t
}

func (t *Template) AddFieldsVerbose(v map[string]verbose.Verbose) *Template {
	t.fieldsVerbose.AddWithMap(v)

	return t
}

func (t *Template) FieldRender(key string, value any) (any, error) {
	field, ok := t.fields[key]
	if ok {
		return field.render(value)
	} else {
		if t.fieldsVerbose != nil {
			return t.fieldsVerbose.ExecuteKey(key, value), nil
		}
	}

	return value, nil
}
