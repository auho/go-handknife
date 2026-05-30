package prompt

import "github.com/manifoldco/promptui"

func NewSelect[T int | string](title string, items []T) (T, error) {
	p := promptui.Select{
		Label: title,
		Items: items,
		Size:  6,
	}

	i, _, err := p.Run()
	if err != nil {
		return *new(T), err
	}

	return items[i], nil
}

func NewSelectWithValue[T int | string](title string, value T, items []T) (T, error) {
	if value == *new(T) {
		_nv, err := NewSelect(title, items)
		if err != nil {
			return *new(T), err
		}

		value = _nv
	}

	return value, nil
}
