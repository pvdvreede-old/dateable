package tests

import "github.com/robfig/revel"

type AppTest struct {
	revel.TestSuite
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}
