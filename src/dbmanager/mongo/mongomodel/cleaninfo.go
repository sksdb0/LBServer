package mongomodel

import (
	"time"
	"util"
)

type CleanInfoModel struct {
	View                 *DailyCleanInfoView
	StartTypemap         map[int]int
	FinishTypemap        map[int]int
	Interpolate_area     *util.Interpolate
	Interpolate_duration *util.Interpolate
	ComplateTypemap      map[int]int
	CleanTimesmap        map[int]int
}

func NewCleanInfoModel(date time.Time) *CleanInfoModel {
	model := CleanInfoModel{
		View:            newDailyCleanInfoView(date),
		StartTypemap:    make(map[int]int),
		FinishTypemap:   make(map[int]int),
		ComplateTypemap: make(map[int]int),
		CleanTimesmap:   make(map[int]int),
	}

	model.StartTypemap[0] = 0
	model.StartTypemap[10] = 1
	model.StartTypemap[20] = 2
	model.StartTypemap[30] = 3

	model.FinishTypemap[0] = 0
	model.FinishTypemap[10] = 1
	model.FinishTypemap[20] = 2
	model.FinishTypemap[30] = 3
	model.FinishTypemap[40] = 4
	model.FinishTypemap[50] = 5
	model.FinishTypemap[60] = 6
	model.FinishTypemap[70] = 7
	model.FinishTypemap[80] = 8
	model.FinishTypemap[90] = 9
	model.FinishTypemap[100] = 10
	model.FinishTypemap[110] = 11
	model.FinishTypemap[120] = 12
	model.FinishTypemap[130] = 13
	model.FinishTypemap[140] = 14
	model.FinishTypemap[150] = 15
	model.FinishTypemap[160] = 16
	model.FinishTypemap[170] = 17
	model.FinishTypemap[180] = 18
	model.FinishTypemap[190] = 19

	model.Interpolate_area = util.NewInterpolate(0, 140, 5)
	model.Interpolate_area.AddRange(0, 5)
	model.Interpolate_area.AddRange(5, 10)
	model.Interpolate_area.AddRange(10, 20)
	model.Interpolate_area.AddRange(20, 40)
	model.Interpolate_area.AddRange(40, 60)
	model.Interpolate_area.AddRange(60, 80)
	model.Interpolate_area.AddRange(80, 100)
	model.Interpolate_area.AddRange(100, 120)
	model.Interpolate_area.AddRange(120, 140)

	model.Interpolate_duration = util.NewInterpolate(0, 80, 10)
	model.Interpolate_duration.AddRange(0, 10)
	model.Interpolate_duration.AddRange(10, 20)
	model.Interpolate_duration.AddRange(20, 30)
	model.Interpolate_duration.AddRange(30, 40)
	model.Interpolate_duration.AddRange(40, 50)
	model.Interpolate_duration.AddRange(50, 60)
	model.Interpolate_duration.AddRange(60, 80)

	return &model
}

type DailyCleanInfoView struct {
	Date               time.Time
	ViewDataStartType  []item
	ViewDataFinishType []item
	ViewDataArea       []item
	ViewDataTimes      []item
	ViewDataDuration   []item
	ViewDataComplate   []item
}

func newDailyCleanInfoView(date time.Time) *DailyCleanInfoView {
	view := DailyCleanInfoView{
		Date: date,
	}

	view.ViewDataStartType = append(view.ViewDataStartType,
		item{Key: "Unknown", Value: 0},
		item{Key: "TimeSchedule", Value: 0},
		item{Key: "ManualStart", Value: 0},
		item{Key: "AppStart", Value: 0})

	view.ViewDataFinishType = append(view.ViewDataFinishType,
		item{Key: "Unknown", Value: 0},
		item{Key: "Complete", Value: 0},
		item{Key: "ManuallyStop", Value: 0},
		item{Key: "NoDisTurbResume", Value: 0},
		item{Key: "FailResume", Value: 0},
		item{Key: "FailResumeII", Value: 0},
		item{Key: "ManuallyShutdown", Value: 0},
		item{Key: "ResetAp", Value: 0},
		item{Key: "BeginLocalClean", Value: 0},
		item{Key: "LowPowerBackToDockFail", Value: 0},
		item{Key: "LowPowerStart", Value: 0},
		item{Key: "SleepTimeout", Value: 0},
		item{Key: "ChangeToSpot", Value: 0},
		item{Key: "DirectlyConnectToDock", Value: 0},
		item{Key: "TimeScheduledCleanInterrupt", Value: 0},
		item{Key: "BatteryVCT", Value: 0},
		item{Key: "UpdateFirmware", Value: 0},
		item{Key: "BatteryForceShutdown", Value: 0},
		item{Key: "PrepareShutdownTimeout", Value: 0},
		item{Key: "CleanFinishType_AppStopCommand", Value: 0})

	view.ViewDataArea = append(view.ViewDataArea,
		item{Key: "0~5m²", Value: 0},
		item{Key: "5~10m²", Value: 0},
		item{Key: "10~20m²", Value: 0},
		item{Key: "20~40m²", Value: 0},
		item{Key: "40~60m²", Value: 0},
		item{Key: "60~80m²", Value: 0},
		item{Key: "80~100m²", Value: 0},
		item{Key: "100~120m²", Value: 0},
		item{Key: "120~140m²", Value: 0},
		item{Key: "140m²以上", Value: 0})

	view.ViewDataDuration = append(view.ViewDataDuration,
		item{Key: "0~10分钟", Value: 0},
		item{Key: "10~20分钟", Value: 0},
		item{Key: "20~30分钟", Value: 0},
		item{Key: "30~40分钟", Value: 0},
		item{Key: "40~50分钟", Value: 0},
		item{Key: "50~60分钟", Value: 0},
		item{Key: "60~80分钟", Value: 0},
		item{Key: "80分钟以上", Value: 0})

	view.ViewDataTimes = append(view.ViewDataTimes,
		item{Key: "1次", Value: 0},
		item{Key: "2次", Value: 0},
		item{Key: "3次", Value: 0},
		item{Key: "4次", Value: 0},
		item{Key: "5次", Value: 0},
		item{Key: "5次以上", Value: 0})

	view.ViewDataComplate = append(view.ViewDataComplate,
		item{Key: "用户手动干预", Value: 0},
		item{Key: "清扫自动完成", Value: 0})

	return &view
}
