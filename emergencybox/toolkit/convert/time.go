package convert

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func TimestampAnyToString(ts any) (string, error) {
	t, err := TimestampToInt64(ts)
	if err != nil {
		return "", err
	}

	s := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	s = fmt.Sprintf("%d [%s]", t, s)

	return s, nil
}

func TimestampToInt64(ts any) (int64, error) {
	var err error
	t := int64(0)
	switch _ts := ts.(type) {
	case string:
		t, err = strconv.ParseInt(_ts, 10, 64)
	case float64:
		t = int64(_ts)
	case int:
		t = int64(_ts)
	case int32:
		t = int64(_ts)
	case uint32:
		t = int64(_ts)
	case time.Time:
		t = _ts.Unix()
	default:
		err = errors.New(fmt.Sprintf("timestamp is error[%T %v]", _ts, _ts))
	}

	return t, err
}
