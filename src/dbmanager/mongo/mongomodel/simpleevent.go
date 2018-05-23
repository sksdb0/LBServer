package mongomodel

import (
	"time"
)

type SimpleEventModel struct {
	View    *DailySimpleEventView
	Typemap map[int]int
}

func NewSimpleEventModel(date time.Time) *SimpleEventModel {
	model := SimpleEventModel{
		View:    newDailySimpleEventView(date),
		Typemap: make(map[int]int),
	}

	return &model
}

type DailySimpleEventView struct {
	Date     time.Time
	ViewData []item
}

func newDailySimpleEventView(date time.Time) *DailySimpleEventView {
	view := DailySimpleEventView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "Spot", Value: 0})

	return &view
}
