package verbose

import (
	"fmt"
	"time"

	"github.com/auho/go-handknife/emergencybox/toolkit/convert"
)

func TimestampCompareNow(v any) any {
	nu := time.Now().Unix()
	_vu, err := convert.TimestampToInt64(v)
	if err != nil {
		return v
	}

	_v1 := ""
	if _vu < nu {
		_v1 = "< now"
	} else if _vu > nu {
		_v1 = "> now"
	}

	return fmt.Sprintf("%v %v", TimestampToString(v), _v1)
}

func TimestampToString(v any) any {
	_v, err := convert.TimestampAnyToString(v)
	if err != nil {
		return v
	} else {
		return _v
	}
}
