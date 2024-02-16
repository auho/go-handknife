package convert

import (
	"regexp"
	"strconv"
)

func ContainsUnicodeToZh(s string) string {
	re := regexp.MustCompile(`(?i)\\u[0-9a-f]{4}`)
	return re.ReplaceAllStringFunc(s, func(s string) string {
		s1, err := strconv.Unquote(`"` + s + `"`)
		if err != nil {
			return s
		} else {
			return s1
		}
	})
}
