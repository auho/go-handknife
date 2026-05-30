package prompt

import "github.com/xuri/excelize/v2"

type Excel struct {
	File      *File
	ExcelFile *excelize.File
}

func (e *Excel) Close() error {
	return e.ExcelFile.Close()
}
