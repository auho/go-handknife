package output

import "unicode/utf8"

// stringFormatLen
// adjust the display length of strings containing chinese characters
func stringFormatLen(s string, maxLen int) int {
	_, cnLen, _ := stringLen(s)
	if cnLen > 0 {
		return maxLen - cnLen
	} else {
		return maxLen
	}
}

// stringLen
// en length
// cn length
// max length
func stringLen(s string) (int, int, int) {
	_utf8Len := utf8.RuneCountInString(s)
	_len := len(s)

	_cnLen := (_len - _utf8Len) / 2
	_enLen := _utf8Len - _cnLen

	return _enLen, _cnLen, _enLen + _cnLen*2
}
