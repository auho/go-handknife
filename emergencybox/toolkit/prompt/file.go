package prompt

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/xuri/excelize/v2"
)

type File struct {
	Path string
	De   fs.DirEntry
}

func (f *File) ToExcel() (*Excel, error) {
	excelFile, err := excelize.OpenFile(f.Path)
	if err != nil {
		return nil, err
	}

	return &Excel{
		File:      f,
		ExcelFile: excelFile,
	}, nil
}

func NewFileBySelectInDir(d string, title string) (*File, error) {
	dirInfo, err := os.Stat(d)
	if err != nil {
		return nil, err
	}

	if !dirInfo.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is not dir", d))
	}

	var files []File

	err = filepath.WalkDir(d, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			files = append(files, File{De: d, Path: path})
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walk dir %w", err)
	}

	var optionsFileName []string
	for _, f := range files {
		optionsFileName = append(optionsFileName, f.De.Name())
	}

	p := promptui.Select{
		Label: title,
		Items: optionsFileName,
		Size:  6,
	}

	i, _, err := p.Run()
	if err != nil {
		return nil, err
	}

	return &(files[i]), nil
}
