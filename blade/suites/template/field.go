package template

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
)

type Field struct {
	name      string
	pipelines Pipelines
	validator func(x any) (any, error)
	detailed  func(x any) (any, error)
}

func (f *Field) AddPipeline(p string) *Field {
	f.pipelines = append(f.pipelines, Pipeline{p})

	return f
}

func (f *Field) AddValidator(fn func(x any) (any, error)) *Field {
	f.validator = fn

	return f
}

func (f *Field) AddDetailed(fn func(x any) (any, error)) *Field {
	f.detailed = fn

	return f
}

func (f *Field) execPipeline(v any) (any, error) {
	return f.pipelines.render(v)
}

func (f *Field) execValidate(v any) (any, error) {
	return f.validator(v)
}

func (f *Field) execDetailded(v any) (any, error) {
	return f.detailed(v)
}

func (f *Field) render(v any) (any, error) {
	if f.validator != nil {
		_, vErr := f.execValidate(v)
		if vErr != nil {
			return fmt.Errorf("%v: %s", v, text.FgRed.Sprintf("%v", vErr)), nil
		}
	}

	if f.detailed != nil {
		detailed, err := f.execDetailded(v)
		if err != nil {
			return fmt.Errorf("%v: %s", v, text.FgRed.Sprintf("%v", err)), nil
		}

		return detailed, nil
	}

	p, err := f.execPipeline(v)
	if err != nil {
		return v, fmt.Errorf("pipeline: %w", err)
	}

	return p, nil
}
