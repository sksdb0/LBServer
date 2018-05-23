package mongomodel

import (
	"time"
)

type ShutdownModel struct {
	View    *DailyShutDownView
	Typemap map[int]int
}

func NewShutdownModel(date time.Time) *ShutdownModel {
	model := ShutdownModel{
		View:    newDailyShutDownView(date),
		Typemap: make(map[int]int),
	}

	model.Typemap[0] = 0
	model.Typemap[10] = 1
	model.Typemap[20] = 2
	model.Typemap[30] = 3
	model.Typemap[40] = 4
	model.Typemap[50] = 5
	model.Typemap[60] = 6
	model.Typemap[70] = 7
	model.Typemap[80] = 8
	model.Typemap[90] = 9
	model.Typemap[100] = 10
	model.Typemap[110] = 11
	model.Typemap[120] = 12
	model.Typemap[130] = 13
	model.Typemap[140] = 14

	return &model
}

type DailyShutDownView struct {
	Date     time.Time
	ViewData []item
}

func newDailyShutDownView(date time.Time) *DailyShutDownView {
	view := DailyShutDownView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "Unknown", Value: 0},
		item{Key: "StartButtonPressHold", Value: 0},
		item{Key: "ResetAp", Value: 0},
		item{Key: "BatteryVCT", Value: 0},
		item{Key: "UpdateFirmware", Value: 0},
		item{Key: "BatteryForceShutdown", Value: 0},
		item{Key: "ScheduledReset", Value: 0},
		item{Key: "SleepTimeout", Value: 0},
		item{Key: "BatteryFault", Value: 0},
		item{Key: "PrepareShutdownTimeout", Value: 0},
		item{Key: "DockDisconnectedAccident", Value: 0},
		item{Key: "McuProxyUnableToStart", Value: 0},
		item{Key: "DockDisconnectedPower", Value: 0},
		item{Key: "MTShutdownTimeout", Value: 0},
		item{Key: "MidnightResetTimeout", Value: 0})

	return &view
}
