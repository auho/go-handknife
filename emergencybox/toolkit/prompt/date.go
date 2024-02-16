package prompt

import (
	"errors"
	"regexp"
)

func NewDateWithValue(title, value string, validate func(string) error) (string, error) {
	if validate == nil {
		validate = func(s string) error {
			re := regexp.MustCompile(`^\d{4}[-/]?\d{2}[-/]?\d{2}$`)

			if re.MatchString(s) {
				return nil
			} else {
				return errors.New("invalid date format")
			}
		}
	}

	var err error
	if value == "" {
		value, err = NewString(title, validate)
		if err != nil {
			return "", err
		}
	} else {
		err = validate(value)
		if err != nil {
			return "", err
		}
	}

	return value, nil
}

func NewDate(title string, validate func(string) error) (string, error) {
	return NewDateWithValue(title, "", validate)
}

func NewDateTimeWithValue(title, value string, validate func(string) error) (string, error) {
	if validate == nil {
		validate = func(s string) error {
			re := regexp.MustCompile(`^\d{4}[-/]?\d{2}[-/]?\d{2} \d{2}:\d{2}:\d{2}$`)
			if re.MatchString(s) {
				return nil
			} else {
				return errors.New("invalid date time format")
			}
		}
	}

	var err error
	if value == "" {
		value, err = NewString(title, validate)
		if err != nil {
			return "", err
		}
	} else {
		err = validate(value)
		if err != nil {
			return "", err
		}
	}

	return value, nil
}

func NewDateTime(title string, validate func(string) error) (string, error) {
	return NewDateTimeWithValue(title, "", validate)
}
