package export

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/xuri/excelize/v2"
)

// FileSliceMap
// excel file
type FileSliceMap struct {
	Filename    string
	FilePath    string
	AbsFilePath string
	AbsLogDir   string
}

func NewResultSliceMap(fileName string) *FileSliceMap {
	currentDir, _ := os.Getwd()

	// save file
	re := regexp.MustCompile(`[<>:"'/\\|?*!\s]`)
	fileName = re.ReplaceAllString(fileName, "_") + ".xlsx"

	filePath := fmt.Sprintf("output/xlsx/%s", fileName)
	return &FileSliceMap{
		Filename:    fileName,
		FilePath:    filePath,
		AbsFilePath: path.Join(currentDir, filePath),
		AbsLogDir:   path.Join(currentDir, "log"),
	}
}

func (r *FileSliceMap) Log(logFileName string) error {
	body := r.Filename + "\n" + r.AbsFilePath
	err := os.WriteFile(path.Join(r.AbsLogDir, "xlsx_"+logFileName), []byte(body), 0644)
	if err != nil {
		return err
	}

	return nil
}

// DataSliceMap
// data of slice map
type DataSliceMap struct {
	IsActive  bool
	SheetName string
	Items     []map[string]any
	Names     []string
	Titles    map[string]string
}

type sheetSliceMap struct {
	data      DataSliceMap
	sheetName string
	isActive  bool
}

// FromSliceMap
// export data of slice map to excel
type FromSliceMap struct {
	fileName string
	fileMark string

	sheets    []sheetSliceMap
	excelFile *excelize.File
}

func NewFromSliceMap(fileName, fileMark string) *FromSliceMap {
	return &FromSliceMap{
		fileName: fileName,
		fileMark: fileMark,
	}
}

func (f *FromSliceMap) Save() (FileSliceMap, error) {
	var ret FileSliceMap

	f.excelFile = excelize.NewFile()
	// add sheets
	err := f.sheetsToExcel()
	if err != nil {
		return ret, fmt.Errorf("add sheets: %w", err)
	}

	err = f.excelFile.DeleteSheet("Sheet1")
	if err != nil {
		return ret, fmt.Errorf("DeleteSheet1 failed: %v", err)
	}

	// save file
	ret = *NewResultSliceMap(f.fileName)
	if err = f.excelFile.SaveAs(ret.FilePath); err != nil {
		return ret, fmt.Errorf("SaveAs[%s] %w", ret.FilePath, err)
	}

	if err = f.excelFile.Close(); err != nil {
		return ret, fmt.Errorf("close[%s] %w", ret.FilePath, err)
	}

	err = ret.Log(f.fileMark)
	if err != nil {
		return ret, fmt.Errorf("log[%s] %w", ret.FilePath, err)
	}

	return ret, nil
}

func (f *FromSliceMap) AddData(data DataSliceMap) *FromSliceMap {
	return f.AddSheet(data, data.SheetName, data.IsActive)
}

// AddSheet
//
// string: excel excelFile path
func (f *FromSliceMap) AddSheet(data DataSliceMap, sheetName string, isActive bool) *FromSliceMap {
	f.sheets = append(f.sheets, sheetSliceMap{
		data:      data,
		sheetName: sheetName,
		isActive:  isActive,
	})

	return f
}

func (f *FromSliceMap) sheetsToExcel() error {
	for _, sheet := range f.sheets {
		err := f.sheetToExcel(sheet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FromSliceMap) sheetToExcel(sheet sheetSliceMap) error {
	_sheetIndex, err := f.excelFile.NewSheet(sheet.sheetName)
	if err != nil {
		return fmt.Errorf("NewSheet %w", err)
	}

	// header
	header := make(map[string]any)
	for _name, _title := range sheet.data.Titles {
		header[_name] = _title
	}

	sheet.data.Items = append([]map[string]any{header}, sheet.data.Items...)

	// data
	var cell string
	for i, _item := range sheet.data.Items {
		cell, err = excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			return fmt.Errorf("CoordinatesToCellName %w", err)
		}

		var _row []any
		for _, _name := range sheet.data.Names {
			_row = append(_row, _item[_name])
		}

		err = f.excelFile.SetSheetRow(sheet.sheetName, cell, &_row)
		if err != nil {
			return fmt.Errorf("SetSheetRow %w", err)
		}
	}

	if sheet.isActive {
		f.excelFile.SetActiveSheet(_sheetIndex)
	}

	return nil
}
