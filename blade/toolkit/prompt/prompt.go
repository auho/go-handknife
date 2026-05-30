package prompt

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func NewString(title string, validate func(string) error) (string, error) {
	result, err := run(title)
	if err != nil {
		return "", err
	}

	if validate != nil {
		err = validate(result)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

func NewStringWithValue(title string, value string, validate func(string) error) (string, error) {
	if value == "" {
		_nv, err := NewString(title, validate)
		if err != nil {
			return "", err
		}

		value = _nv
	}

	return value, nil
}

func NewInt64(title string, validate func(int64) error) (int64, error) {
	result, err := run(title)
	if err != nil {
		return 0, err
	}

	i64, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	if validate != nil {
		err = validate(i64)
		if err != nil {
			return 0, err
		}
	}

	return i64, nil
}

func NewInt(title string, validate func(int) error) (int, error) {
	result, err := run(title)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(result)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	if validate != nil {
		err = validate(i)
		if err != nil {
			return 0, err
		}
	}

	return i, nil
}

func NewIntWithValue(title string, value int, validate func(int) error) (int, error) {
	if value == 0 {
		_nv, err := NewInt(title, validate)
		if err != nil {
			return 0, err
		}

		value = _nv
	}

	return value, nil
}

func NewFloat64(title string, validate func(float64) error) (float64, error) {
	result, err := run(title)
	if err != nil {
		return 0, err
	}

	f64, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	if validate != nil {
		err = validate(f64)
		if err != nil {
			return 0, err
		}
	}

	return f64, nil
}

func NewFloat64WithValue(title string, value float64, validate func(float64) error) (float64, error) {
	if value == 0 {
		_nv, err := NewFloat64(title, validate)
		if err != nil {
			return 0, err
		}

		value = _nv
	}

	return value, nil
}

func run(title string) (string, error) {
	prompt := promptui.Prompt{
		Label: title,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %w", err)
	}

	return result, nil
}
