package strings

import "strings"

func ToUnderlineNaming(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	sb := []byte(s)

	var ssb []byte
	for _, b := range sb {
		if b >= 65 && b <= 90 {
			ssb = append(ssb, 95, b+32)
		} else {
			ssb = append(ssb, b)
		}
	}

	return string(ssb)
}

func ToHumpNaming(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	ss := strings.Split(s, "_")
	for k, split := range ss {
		if k > 0 {
			ss[k] = strings.ToUpper(split[0:1]) + split[1:]
		}
	}

	return strings.Join(ss, "")
}

func ToUpperFirstChar(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
