package tests

import "github.com/robfig/revel"

type DateTest struct {
  revel.TestSuite
}

func (t DateTest) TestDateWithWeekday() {
  t.Get("/date/2013-03-04")
  t.AssertOk()
  t.AssertContentType("application/json")
  t.AssertContains("\"Weekend\": false")
  t.AssertContains("\"Date\": \"2013-03-04\"")
}

func (t DateTest) TestDateWithWeekend() {
  t.Get("/date/2013-03-02")
  t.AssertOk()
  t.AssertContentType("application/json")
  t.AssertContains("\"Weekend\": true")
  t.AssertContains("\"Date\": \"2013-03-02\"")
}
