package mongomodel

import (
	"time"
	"util"
)

type RemoteModel struct {
	View                 *DailyRemoteView
	Typemap              map[int]int
	Interpolate_duration *util.Interpolate
}

func NewRemoteModel(date time.Time) *RemoteModel {
	model := RemoteModel{
		View:    newDailyRemoteView(date),
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
	model.Typemap[150] = 15
	model.Typemap[160] = 16
	model.Typemap[170] = 17

	model.Interpolate_duration = util.NewInterpolate(0, 600, 15)
	model.Interpolate_duration.AddRange(0, 15)
	model.Interpolate_duration.AddRange(15, 30)
	model.Interpolate_duration.AddRange(30, 60)
	model.Interpolate_duration.AddRange(60, 120)
	model.Interpolate_duration.AddRange(120, 180)
	model.Interpolate_duration.AddRange(180, 300)
	model.Interpolate_duration.AddRange(300, 600)

	return &model
}

type DailyRemoteView struct {
	Date              time.Time
	ViewDataRemote    []item
	ViewDataDuration  []item
	ViewDataPreState  []item
	ViewDataNextState []item
}

func newDailyRemoteView(date time.Time) *DailyRemoteView {
	view := DailyRemoteView{
		Date: date,
	}

	view.ViewDataRemote = append(view.ViewDataRemote,
		item{Key: "Remote", Value: 0})

	view.ViewDataPreState = append(view.ViewDataPreState,
		item{Key: "Unknown", Value: 0},
		item{Key: "Invalid", Value: 0},
		item{Key: "Initial", Value: 0},
		item{Key: "Sleep", Value: 0},
		item{Key: "Standby", Value: 0},
		item{Key: "Remote", Value: 0},
		item{Key: "Cleaning", Value: 0},
		item{Key: "BackToDock", Value: 0},
		item{Key: "SearchForDock", Value: 0},
		item{Key: "Charging", Value: 0},
		item{Key: "ChargingError", Value: 0},
		item{Key: "Pause", Value: 0},
		item{Key: "Spot", Value: 0},
		item{Key: "Malfunctioning", Value: 0},
		item{Key: "PrepareShutdown", Value: 0},
		item{Key: "Updating", Value: 0},
		item{Key: "RubToDock", Value: 0},
		item{Key: "MobilityTest", Value: 0})

	view.ViewDataNextState = append(view.ViewDataNextState,
		item{Key: "Unknown", Value: 0},
		item{Key: "Invalid", Value: 0},
		item{Key: "Initial", Value: 0},
		item{Key: "Sleep", Value: 0},
		item{Key: "Standby", Value: 0},
		item{Key: "Remote", Value: 0},
		item{Key: "Cleaning", Value: 0},
		item{Key: "BackToDock", Value: 0},
		item{Key: "SearchForDock", Value: 0},
		item{Key: "Charging", Value: 0},
		item{Key: "ChargingError", Value: 0},
		item{Key: "Pause", Value: 0},
		item{Key: "Spot", Value: 0},
		item{Key: "Malfunctioning", Value: 0},
		item{Key: "PrepareShutdown", Value: 0},
		item{Key: "Updating", Value: 0},
		item{Key: "RubToDock", Value: 0},
		item{Key: "MobilityTest", Value: 0})

	view.ViewDataDuration = append(view.ViewDataDuration,
		item{Key: "0~15s", Value: 0},
		item{Key: "15~30s", Value: 0},
		item{Key: "30~60s", Value: 0},
		item{Key: "1~2m", Value: 0},
		item{Key: "2~3m", Value: 0},
		item{Key: "3~5m", Value: 0},
		item{Key: "5~10m", Value: 0},
		item{Key: "10m以上", Value: 0})

	return &view
}
