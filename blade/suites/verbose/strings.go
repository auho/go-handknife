package verbose

import (
	"github.com/auho/go-handknife/blade/toolkit/strutil"
	strings2 "github.com/auho/go-toolkit/farmtools/convert/types/strings"
)

func ToUnderlineNaming(v any) any {
	vs := v.(string)
	return strutil.ToUnderlineNaming(vs)
}

func Truncate(v any, l int) any {
	_vs, err := strings2.FromAny(v)
	if err != nil {
		return v
	}

	_len := len(_vs)
	if _len <= l {
		return _vs
	}

	return _vs[0:l] + "..."
}
