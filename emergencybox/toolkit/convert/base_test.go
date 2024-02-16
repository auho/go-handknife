package convert

import "testing"

func TestBase(t *testing.T) {
	// 1nje4y1nje581nje591nje5a1nje5b1nje5c1nje5e
	// 100002130,100002140,100002141,100002142,100002143,100002144,100002146

	nums := []int{100002130, 100002140, 100002141, 100002142, 100002143, 100002144, 100002146}
	nums36 := []string{"1nje4y", "1nje58", "1nje59", "1nje5a", "1nje5b", "1nje5c", "1nje5e"}

	for k, num := range nums {
		num36 := DecimalToBHex(num, 36)
		if num36 != nums36[k] {
			t.Error("36 error", num, num36, nums36[k])
		}

		newNum := BHex2Decimal(num36, 36)
		if newNum != num {
			t.Error("10 error", num, num36, newNum)
		}
	}
}
