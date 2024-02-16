package prompt

import "github.com/manifoldco/promptui"

type CustomSelect[T int | string] struct {
	Value T
	Title string
}

func NewCustomSelect[T int | string](title string, items []CustomSelect[T]) (T, error) {
	//var titles []string
	//for _, item := range items {
	//	titles = append(titles, item.Title)
	//}

	p := promptui.Select{
		Label: title,
		Items: items,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F336  {{ .Value | cyan }} ({{ .Title | red }})",
			Inactive: "  {{ .Value | cyan }} ({{ .Title | red }})",
			Selected: "\U0001F336  {{ .Title | red | cyan }}",
			Details: `
{{ "Value:" | faint }}	{{ .Value }}
{{ "Title:" | faint }}	{{ .Title }}`,
		},
		Size: 6,
	}

	i, _, err := p.Run()
	if err != nil {
		return *new(T), err
	}

	return items[i].Value, nil
}

func NewCustomSelectWithValue[T int | string](title string, value T, items []CustomSelect[T]) (T, error) {
	if value == *new(T) {
		_nv, err := NewCustomSelect(title, items)
		if err != nil {
			return *new(T), err
		}

		value = _nv
	}

	return value, nil
}
