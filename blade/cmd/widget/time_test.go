package widget

import (
	"testing"
)

func TestNewTime(t *testing.T) {
	_t := NewTime()

	err := _t.ServerTimeContrast(0)
	if err != nil {
		t.Error(err)
	}
}
