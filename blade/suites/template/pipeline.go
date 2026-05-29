package template

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Pipelines []Pipeline

func (ps Pipelines) Add(s string) {
	ps = append(ps, Pipeline{p: s})
}

func (ps Pipelines) render(v any) (string, error) {
	pipeline := ps.pipelines()
	if pipeline == "" {
		pipeline = "{{.}}"
	} else {
		pipeline = "{{. | " + pipeline + "}}"
	}

	tmpl, err := template.New("").Parse(pipeline)
	if err != nil {
		return "", fmt.Errorf("template parse[%s]: %w", pipeline, err)
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, v)
	if err != nil {
		return "", fmt.Errorf("template execute[%s: %w", pipeline, err)
	}

	return b.String(), nil
}

func (ps Pipelines) pipelines() string {
	var ss []string
	for _, p := range ps {
		ss = append(ss, p.pipeline())
	}

	return strings.Join(ss, " | ")
}

type Pipeline struct {
	p string
}

func (p *Pipeline) pipeline() string {
	return p.p
}
