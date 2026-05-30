package suites

import (
	"github.com/auho/go-handknife/blade/suites/output"
	"github.com/auho/go-handknife/blade/suites/template"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Post struct {
	template *template.Template
	output   *output.Output
	extra    func() error
}

func NewPost() *Post {
	return &Post{
		template: template.NewTemplate(),
		output:   output.NewOutputDefault(),
	}
}

func (p *Post) SettingTemplate(fn func(t *template.Template)) *Post {
	fn(p.template)

	return p
}

func (p *Post) SettingOutput(fn func(w table.Writer)) *Post {
	p.output.Setting(fn)

	return p
}

func (p *Post) AddExtra(fn func() error) *Post {
	p.extra = fn

	return p
}

func (p *Post) ExtraExec() (bool, error) {
	if p.extra != nil {
		return true, p.extra()
	}

	return false, nil
}

func (p *Post) render(sm []map[string]any) (string, error) {
	var err error
	sm, err = p.renderSliceMap(sm)
	if err != nil {
		return "", err
	}

	body, err := p.output.Output(sm, p.template)
	if err != nil {
		return "", err
	}

	return body, nil
}

func (p *Post) renderSliceMap(sm []map[string]any) ([]map[string]any, error) {
	var nsm []map[string]any
	for _, m := range sm {
		nm, err := p.renderMap(m)
		if err != nil {
			return nil, err
		}

		nsm = append(nsm, nm)
	}

	return nsm, nil
}

func (p *Post) renderMap(m map[string]any) (map[string]any, error) {
	nm := make(map[string]any, len(m))
	for k, v := range m {
		val, err := p.template.FieldRender(k, v)
		if err != nil {
			return nil, err
		}

		nm[k] = val
	}

	return nm, nil
}
