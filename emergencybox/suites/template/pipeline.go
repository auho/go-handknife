package template

import "strings"

type Pipelines []Pipeline

func (ps Pipelines) Add(s string) {
	ps = append(ps, Pipeline{p: s})
}

func (ps Pipelines) Render() string {
	var ss []string
	for _, p := range ps {
		ss = append(ss, p.Pipeline())
	}

	return strings.Join(ss, " | ")
}

type Pipeline struct {
	p string
}

func (p *Pipeline) Pipeline() string {
	return p.p
}
