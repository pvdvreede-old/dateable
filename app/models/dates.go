package models

import (
	"time"
)

const dayInNano = time.Hour * 24

type Dates struct {
	TotalDays int
	Results   []Date
}

func (d *Dates) GetTotalDays() int {
	return len(d.Results)
}

func NewDates(from, to string) Dates {
	f := GetDate(from)
	t := GetDate(to)
	var dates Dates
	dates.Results = GetDates(f, t)
	dates.TotalDays = dates.GetTotalDays()
	return dates
}

func GetDates(from, to time.Time) []Date {
	cont := true
	var dates []Date
	incDate := from
	for cont {
		if incDate.After(to) {
			cont = false
		} else {
			dates = append(dates, NewDateFromTime(incDate))
			incDate = incDate.Add(dayInNano)
		}
	}
	return dates
}
