package models

import "time"

type Date struct {
  Date string
  Weekend bool
}

func NewDate(dateStr string) Date {
  t := getDate(dateStr)
  date := Date{}
  date.Date = t.Format("2006-01-02")
  date.Weekend = isWeekend(t)
  return date
}

func getDate(dateStr string) time.Time {
  t, _ := time.Parse("2006-01-02", dateStr)
  return t
}

func isWeekend(t time.Time) bool {
  if w := t.Weekday(); w == 0 || w == 6 {
    return true
  }
  return false
}
