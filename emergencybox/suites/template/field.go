package template

type Field struct {
	name      string
	pipelines Pipelines
	validator func(x any, ps Pipelines) (any, Pipelines, error)
}

func (f *Field) AddPipeline(p string) *Field {
	f.pipelines = append(f.pipelines, Pipeline{p})

	return f
}

func (f *Field) AddValidator(fn func(x any, ps Pipelines) (any, Pipelines, error)) *Field {
	f.validator = fn

	return f
}

func (f *Field) Pipelines() Pipelines {
	return f.pipelines
}

func (f *Field) Validate(x any) (any, Pipelines, error) {
	if f.validator == nil {
		return x, f.pipelines, nil
	}

	return f.validator(x, f.pipelines)
}
