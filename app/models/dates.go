package models

import (
	"time"
)

const dayInNano = time.Hour * 24

type Dates []Date

func NewDates(from, to string) Dates {
	f := GetDate(from)
	t := GetDate(to)
	return GetDates(f, t)
}

func GetDates(from, to time.Time) Dates {
	cont := true
	var dates Dates
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
