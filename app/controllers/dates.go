package controllers

import "github.com/robfig/revel"
import "github.com/pvdvreede/dateable/app/models"

type Dates struct {
	*revel.Controller
}

func (c Dates) Index(from, to string) revel.Result {
	dates := models.NewDates(from, to)
	return c.RenderJson(dates)
}
