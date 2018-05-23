package mongomodel

import (
	"time"
)

type DockModel struct {
	View    *DailyDockView
	Typemap map[int]int
}

func NewDockModel(date time.Time) *DockModel {
	errormodel := DockModel{
		View:    newDailyDockView(date),
		Typemap: make(map[int]int),
	}

	errormodel.Typemap[0] = 0
	errormodel.Typemap[10] = 1
	errormodel.Typemap[20] = 2
	errormodel.Typemap[30] = 3
	errormodel.Typemap[40] = 4
	errormodel.Typemap[50] = 5
	errormodel.Typemap[60] = 6

	return &errormodel
}

type DailyDockView struct {
	Date     time.Time
	ViewData []item
}

func newDailyDockView(date time.Time) *DailyDockView {
	view := DailyDockView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "Unknown", Value: 0},
		item{Key: "BackToDockSucc", Value: 0},
		item{Key: "BackToDockTrapAndNearDock", Value: 0},
		item{Key: "BackToOriginSucc", Value: 0},
		item{Key: "BackToOriginFail", Value: 0},
		item{Key: "BackToDockNoPower", Value: 0},
		item{Key: "DockResult_Malfunctioning", Value: 0})

	return &view
}
