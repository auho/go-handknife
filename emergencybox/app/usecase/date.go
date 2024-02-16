package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/auho/go-handknife/emergencybox/app"
	"github.com/auho/go-handknife/emergencybox/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ app.UseCastor = (*DateUseCase)(nil)

type DateUseCase struct {
	StartDate string
	EndDate   string

	StartDateTime string
	EndTDateTime  string

	DateIndices []string
	DateAmount  int

	startTime time.Time
	endTime   time.Time
}

func (duc *DateUseCase) handleDate() error {
	duc.StartDate = strings.TrimSpace(duc.StartDate)
	duc.EndDate = strings.TrimSpace(duc.EndDate)

	duc.StartDateTime = strings.TrimSpace(duc.StartDate) + " 00:00:00"
	duc.EndTDateTime = strings.TrimSpace(duc.EndDate) + " 23:59:59"

	var err error
	// date indices
	duc.startTime, err = time.Parse(time.DateOnly, duc.StartDate)
	if err != nil {
		return err
	}

	duc.endTime, err = time.Parse(time.DateOnly, duc.EndDate)
	if err != nil {
		return err
	}

	duc.DateIndices = duc.IndicesWithFormat(time.DateOnly)
	duc.DateAmount = len(duc.DateIndices)

	return nil
}

func (duc *DateUseCase) IndicesWithFormat(layout string) []string {
	_tTime := duc.startTime
	_endDate := duc.endTime.Format(layout)
	var indices []string
	for {
		_date := _tTime.Format(layout)
		if _date > _endDate {
			break
		}

		indices = append(indices, _date)
		_tTime = _tTime.AddDate(0, 0, 1)
	}

	return indices
}

func (duc *DateUseCase) ArgumentDay() error {
	var err error
	duc.StartDate, err = prompt.NewDateTimeWithValue("start dtae", duc.StartDate, nil)
	if err != nil {
		return err
	}

	duc.EndDate = duc.StartDate

	return duc.handleDate()
}

func (duc *DateUseCase) ArgumentToday() error {
	duc.StartDate = time.Now().Format("2006-01-02")
	duc.EndDate = duc.StartDate

	return duc.handleDate()
}

func (duc *DateUseCase) ArgumentYesterday() error {
	duc.StartDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	duc.EndDate = duc.StartDate

	return duc.handleDate()
}

func (duc *DateUseCase) ArgumentLastMonth() error {
	year, month, _ := time.Now().Date()
	preMonth := month - 1
	preYear := year
	if preMonth == 0 {
		preMonth = 12
		preYear--
	}

	monthFirstDay := time.Date(preYear, preMonth, 1, 0, 0, 0, 0, time.Local)
	nextMonthFirstDay := monthFirstDay.AddDate(0, 1, 0)
	monthLastDay := nextMonthFirstDay.Add(-24 * time.Hour)

	duc.StartDate = monthFirstDay.Format(time.DateOnly)
	duc.EndDate = monthLastDay.Format(time.DateOnly)

	return duc.handleDate()
}

func (duc *DateUseCase) ArgumentLastWeek() error {
	nowTime := time.Now()
	lastMonday := nowTime.AddDate(0, 0, -int(nowTime.Weekday())-6)
	lastSunday := lastMonday.AddDate(0, 0, 6)

	duc.StartDate = lastMonday.Format(time.DateOnly)
	duc.EndDate = lastSunday.Format(time.DateOnly)

	return duc.handleDate()
}

func (duc *DateUseCase) ArgumentAct() error {
	const actLastMonth = "lastMonth"
	const actLastWeek = "lastWeek"
	const actOther = "other"
	var err error
	var act string
	act, err = prompt.NewSelectWithValue("date act", act, []string{actLastMonth, actLastWeek, actOther})
	if err != nil {
		return err
	}

	switch act {
	case actLastMonth:
		return duc.ArgumentLastMonth()
	case actLastWeek:
		return duc.ArgumentLastWeek()
	case actOther:
		return duc.ParseArgument()
	}

	return nil
}

func (duc *DateUseCase) ArgumentActDay() error {
	const actYesterday = "yesterday"
	const actToday = "today"
	const actOther = "other"

	var err error
	var act string
	act, err = prompt.NewSelectWithValue("date act", act, []string{actYesterday, actToday, actOther})
	if err != nil {
		return err
	}

	switch act {
	case actYesterday:
		return duc.ArgumentYesterday()
	case actToday:
		return duc.ArgumentToday()
	case actOther:
		return duc.ArgumentDay()
	}

	return nil
}

func (duc *DateUseCase) ParseArgument() error {
	var err error
	duc.StartDate, err = prompt.NewDateTimeWithValue("start date", duc.StartDate, nil)
	if err != nil {
		return err
	}

	duc.EndDate, err = prompt.NewDateTimeWithValue("end date", duc.EndDate, nil)
	if err != nil {
		return err
	}

	return duc.handleDate()
}

func (duc *DateUseCase) StartFormat(layout string) string {
	return duc.startTime.Format(layout)
}

func (duc *DateUseCase) EndFormat(layout string) string {
	return duc.endTime.Format(layout)
}

func (duc *DateUseCase) CmdArgs() []string {
	if duc.StartDate == "" {
		return nil
	}

	return []string{
		fmt.Sprintf("--start-date %s --end-date %s", duc.StartDate, duc.EndDate),
	}
}

func (duc *DateUseCase) CmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&duc.StartDate, "start-date", "", "start date")
	cmd.Flags().StringVar(&duc.EndDate, "end-date", "", "end date")
}
