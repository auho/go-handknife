package widget

import (
	"testing"
)

func TestNewTime(t *testing.T) {
	_t := NewTime()

	err := _t.ServerTimeContrast()
	if err != nil {
		t.Error(err)
	}
}
