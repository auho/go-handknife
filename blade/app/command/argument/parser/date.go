package parser

import (
	"fmt"
	"strings"
	"time"

	"github.com/auho/go-handknife/blade/toolkit/prompt"
	"github.com/spf13/cobra"
)

var _ Parser = (*Date)(nil)

type Date struct {
	StartDate string // 2016-01-02
	EndDate   string // 2016-01-02

	StartDateTime string // 2016-01-02 00:00:00
	EndDateTime   string // 2016-01-02 23:59:59

	DateIndices []string // []2016-01-02
	DateAmount  int

	startTime time.Time
	endTime   time.Time
}

func (d *Date) ParseDay() error {
	var err error
	d.StartDate, err = prompt.NewDateTimeWithValue("start date", d.StartDate, nil)
	if err != nil {
		return err
	}

	d.EndDate = d.StartDate

	return d.handleDate()
}

func (d *Date) ParseToday() error {
	d.StartDate = time.Now().Format("2006-01-02")
	d.EndDate = d.StartDate

	return d.handleDate()
}

func (d *Date) ParseYesterday() error {
	d.StartDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	d.EndDate = d.StartDate

	return d.handleDate()
}

func (d *Date) ParseLastMonth() error {
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

	d.StartDate = monthFirstDay.Format(time.DateOnly)
	d.EndDate = monthLastDay.Format(time.DateOnly)

	return d.handleDate()
}

func (d *Date) ParseLastWeek() error {
	nowTime := time.Now()
	lastMonday := nowTime.AddDate(0, 0, -int(nowTime.Weekday())-6)
	lastSunday := lastMonday.AddDate(0, 0, 6)

	d.StartDate = lastMonday.Format(time.DateOnly)
	d.EndDate = lastSunday.Format(time.DateOnly)

	return d.handleDate()
}

func (d *Date) ParseAct() error {
	const actLastMonth = "lastMonth"
	const actLastWeek = "lastWeek"
	const actDefault = "default"

	var err error
	var act string
	act, err = prompt.NewSelectWithValue("date act", act, []string{actLastMonth, actLastWeek, actDefault})
	if err != nil {
		return err
	}

	switch act {
	case actLastMonth:
		return d.ParseLastMonth()
	case actLastWeek:
		return d.ParseLastWeek()
	case actDefault:
		return d.Parse()
	}

	return nil
}

func (d *Date) ParseActDay() error {
	const actYesterday = "yesterday"
	const actToday = "today"
	const actDay = "day"

	var err error
	var act string
	act, err = prompt.NewSelectWithValue("date act", act, []string{actYesterday, actToday, actDay})
	if err != nil {
		return err
	}

	switch act {
	case actYesterday:
		return d.ParseYesterday()
	case actToday:
		return d.ParseToday()
	case actDay:
		return d.ParseDay()
	}

	return nil
}

func (d *Date) Flags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&d.StartDate, "start-date", "", "start date")
	cmd.Flags().StringVar(&d.EndDate, "end-date", "", "end date")
}

func (d *Date) Parse() error {
	var err error
	d.StartDate, err = prompt.NewDateTimeWithValue("start date", d.StartDate, nil)
	if err != nil {
		return err
	}

	d.EndDate, err = prompt.NewDateTimeWithValue("end date", d.EndDate, nil)
	if err != nil {
		return err
	}

	return d.handleDate()
}

func (d *Date) ArgsToString() []string {
	if d.StartDate == "" {
		return nil
	}

	return []string{
		fmt.Sprintf("--start-date %s --end-date %s", d.StartDate, d.EndDate),
	}
}

func (d *Date) StartFormat(layout string) string {
	return d.startTime.Format(layout)
}

func (d *Date) EndFormat(layout string) string {
	return d.endTime.Format(layout)
}

func (d *Date) handleDate() error {
	d.StartDate = strings.TrimSpace(d.StartDate)
	d.EndDate = strings.TrimSpace(d.EndDate)

	d.StartDateTime = strings.TrimSpace(d.StartDate) + " 00:00:00"
	d.EndDateTime = strings.TrimSpace(d.EndDate) + " 23:59:59"

	var err error
	// date indices
	d.startTime, err = time.Parse(time.DateOnly, d.StartDate)
	if err != nil {
		return err
	}

	d.endTime, err = time.Parse(time.DateOnly, d.EndDate)
	if err != nil {
		return err
	}

	d.DateIndices = d.indicesWithFormat(time.DateOnly)
	d.DateAmount = len(d.DateIndices)

	return nil
}

func (d *Date) indicesWithFormat(layout string) []string {
	_tTime := d.startTime
	_endDate := d.endTime.Format(layout)
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
