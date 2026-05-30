package verbose

import (
	"fmt"
	"time"

	"github.com/auho/go-handknife/blade/toolkit/convert"
)

func TimeAnyCompareNow(v any) any {
	nu := time.Now().Unix()
	_vu, err := convert.TimeAnyToInt64(v)
	if err != nil {
		return v
	}

	_v1 := ""
	if _vu < nu {
		_v1 = "< now"
	} else if _vu > nu {
		_v1 = "> now"
	}

	return fmt.Sprintf("%v %v", TimeAnyToString(v), _v1)
}

func TimeAnyToString(v any) any {
	_v, err := convert.TimeAnyToString(v)
	if err != nil {
		return v
	}

	return _v
}
