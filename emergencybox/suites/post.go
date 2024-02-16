package suites

import (
	"github.com/auho/go-handknife/emergencybox/suites/template"
)

type Post struct {
	Template *template.Template
	Extra    func() error
}

func NewPost() *Post {
	return &Post{Template: template.NewTemplate()}
}

func (p *Post) AddTemplate(fn func(t *template.Template)) *Post {
	fn(p.Template)

	return p
}

func (p *Post) AddExtra(fn func() error) *Post {
	p.Extra = fn

	return p
}
