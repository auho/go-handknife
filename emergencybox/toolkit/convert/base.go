package convert

import (
	"math"
	"strings"
)

var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

// DecimalToBHex 10 进制数转换   n 表示进制， 16 or 36
func DecimalToBHex(num, n int) string {
	numStr := ""
	for num != 0 {
		yu := num % n
		numStr = string(num2char[yu]) + numStr
		num = num / n
	}

	return numStr
}

// BHex2Decimal 36 进制数转换   n 表示进制， 16 or 36
func BHex2Decimal(str string, n int) int {
	str = strings.ToLower(str)

	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒序
	}

	return int(v)
}
