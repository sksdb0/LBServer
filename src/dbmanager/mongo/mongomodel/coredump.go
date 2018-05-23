package mongomodel

import (
	"time"
)

type CoredumpModel struct {
	View    *DailyCoreDumpView
	Typemap map[int]int
}

func NewCoredumpModel(date time.Time) *CoredumpModel {
	model := CoredumpModel{
		View:    newDailyCoreDumpView(date),
		Typemap: make(map[int]int),
	}

	return &model
}

type DailyCoreDumpView struct {
	Date     time.Time
	ViewData []item
}

func newDailyCoreDumpView(date time.Time) *DailyCoreDumpView {
	view := DailyCoreDumpView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "CoreDumpCounts", Value: 0},
		item{Key: "CoreDumpRobotCounts", Value: 0})
	return &view
}
