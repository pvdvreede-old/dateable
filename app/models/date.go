package models

import "time"

type Date struct {
	Date    string
	Weekend bool
}

func NewDate(dateStr string) Date {
	t := GetDate(dateStr)
	return NewDateFromTime(t)
}

func NewDateFromTime(time time.Time) Date {
	date := Date{}
	date.Date = time.Format("2006-01-02")
	date.Weekend = isWeekend(time)
	return date
}

func GetDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}

func isWeekend(t time.Time) bool {
	if w := t.Weekday(); w == 0 || w == 6 {
		return true
	}
	return false
}
