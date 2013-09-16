package models

import (
	"time"
)

const dayInNano = time.Hour * 24

type Dates struct {
	TotalDays   int
	Weekdays    int
	WeekendDays int
	Results     []Date
}

func (d *Dates) PopulateStats() {
	weekdays := 0
	for _, v := range d.Results {
		if !v.Weekend {
			weekdays++
		}
	}
	d.TotalDays = len(d.Results)
	d.Weekdays = weekdays
	d.WeekendDays = d.TotalDays - d.Weekdays
}

func NewDates(from, to string) Dates {
	f := GetDate(from)
	t := GetDate(to)
	var dates Dates
	dates.Results = GetDates(f, t)
	dates.PopulateStats()
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
