package verbose

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonIndent(v any) any {
	if _vs, ok := v.(string); ok {
		var b bytes.Buffer
		err1 := json.Indent(&b, []byte(_vs), "  ", "  ")
		if err1 != nil {
			return v
		}

		return b.String()
	} else {
		return v
	}
}

func JsonMarshalFields(fields []string) func(any) any {
	return func(v any) any {
		if _vms, ok := v.(string); ok {
			var _vm map[string]any
			err1 := json.Unmarshal([]byte(_vms), &_vm)
			if err1 != nil {
				return ""
			}
			s := ""
			for _, f := range fields {
				s += fmt.Sprintf("%s:%v ", f, _vm[f])
			}

			return s
		} else {
			return ""
		}
	}
}
