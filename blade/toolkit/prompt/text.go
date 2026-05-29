package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewText(title string) (string, error) {
	return NewTextWithSize(title, 8192)
}

func NewTextWithSize(title string, size int) (string, error) {
	fmt.Println(fmt.Sprintf("%s[%d]:", title, size))

	var ss []string
	reader := bufio.NewReaderSize(os.Stdin, size)
	for {
		s, err := reader.ReadString('\n')
		if s == "\n" {
			_, err = os.Stdin.Write([]byte{4})
			if err != nil {
				return "", err
			}

			break
		}

		if err != nil {
			return "", err
		}

		ss = append(ss, s)
	}

	return strings.Trim(strings.Join(ss, ""), "\n"), nil
}

func NewTextFromFileOrInput(title, filename string) (string, error) {
	return NewTextFromFileOrInputWithSize(title, filename, 8092)
}

func NewTextFromFileOrInputWithSize(title, filename string, size int) (string, error) {
	var text string
	var err error
	if filename != "" {
		var _b []byte
		_b, err = os.ReadFile(filename)
		if err != nil {
			return "", err
		}

		text = string(_b)
	} else {
		text, err = NewTextWithSize(title, size)
		if err != nil {
			return "", err
		}
	}

	return text, nil
}
