package mongomodel

import (
	"time"
)

type ErrorModel struct {
	View    *DailyErrorView
	Typemap map[int]int
}

func NewErrorModel(date time.Time) *ErrorModel {
	model := ErrorModel{
		View:    newDailyErrorView(date),
		Typemap: make(map[int]int),
	}

	model.Typemap[255] = 0
	model.Typemap[256] = 1
	model.Typemap[257] = 2
	model.Typemap[258] = 3
	model.Typemap[259] = 4
	model.Typemap[260] = 5
	model.Typemap[261] = 6
	model.Typemap[262] = 7
	model.Typemap[263] = 8
	model.Typemap[1] = 9
	model.Typemap[2] = 10
	model.Typemap[3] = 11
	model.Typemap[4] = 12
	model.Typemap[5] = 13
	model.Typemap[6] = 14
	model.Typemap[7] = 15
	model.Typemap[8] = 16
	model.Typemap[9] = 17
	model.Typemap[10] = 18
	model.Typemap[11] = 19
	model.Typemap[12] = 20
	model.Typemap[13] = 21
	model.Typemap[14] = 22
	model.Typemap[15] = 23
	model.Typemap[16] = 24
	model.Typemap[17] = 25
	model.Typemap[18] = 26
	model.Typemap[19] = 27
	model.Typemap[20] = 28

	return &model
}

type DailyErrorView struct {
	Date     time.Time
	ViewData []item
}

func newDailyErrorView(date time.Time) *DailyErrorView {
	view := DailyErrorView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "InternalError", Value: 0},
		item{Key: "CompassRecover", Value: 0},
		item{Key: "MBrushSC", Value: 0},
		item{Key: "MBrushOC", Value: 0},
		item{Key: "LWheelSC", Value: 0},
		item{Key: "LWheelOC", Value: 0},
		item{Key: "RWheelSC", Value: 0},
		item{Key: "RWheelOC", Value: 0},
		item{Key: "FanOC", Value: 0},
		item{Key: "Error1", Value: 0},
		item{Key: "Error2", Value: 0},
		item{Key: "Error3", Value: 0},
		item{Key: "Error4", Value: 0},
		item{Key: "Error5", Value: 0},
		item{Key: "Error6", Value: 0},
		item{Key: "Error7", Value: 0},
		item{Key: "Error8", Value: 0},
		item{Key: "Error9", Value: 0},
		item{Key: "Error10", Value: 0},
		item{Key: "Error11", Value: 0},
		item{Key: "Error12", Value: 0},
		item{Key: "Error13", Value: 0},
		item{Key: "Error14", Value: 0},
		item{Key: "Error15", Value: 0},
		item{Key: "Error16", Value: 0},
		item{Key: "Error17", Value: 0},
		item{Key: "Error18", Value: 0},
		item{Key: "Error19", Value: 0},
		item{Key: "Error20", Value: 0})

	return &view
}
