package datetime

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

func StartDateTime(fl validator.FieldLevel) bool {
	return fieldDateTime(fl).Before(time.Now())
}

func EndDateTime(fl validator.FieldLevel) bool {
	return fieldDateTime(fl).After(time.Now())
}

func fieldDateTime(fl validator.FieldLevel) time.Time {
	var t time.Time
	switch fv := fl.Field().Interface().(type) {
	case time.Time:
		t = fv
	case string:
		loc, _ := time.LoadLocation("Local")
		t, _ = time.ParseInLocation("2006-01-02 15:04:05", fv, loc)
	case int:
		t = time.Unix(int64(fv), 0)
	case float64:
		t = time.Unix(int64(fv), 0)
	default:
		log.Fatal("date time error")
	}

	return t
}
