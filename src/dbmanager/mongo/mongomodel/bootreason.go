package mongomodel

import (
	"time"
)

type BootReasonModel struct {
	View    *DailyBootReasonView
	Typemap map[int]int
}

func NewBootReasonModel(date time.Time) *BootReasonModel {
	model := BootReasonModel{
		View:    newDailyBootReasonView(date),
		Typemap: make(map[int]int),
	}

	model.Typemap[4] = 0
	model.Typemap[6] = 1
	model.Typemap[10] = 2

	return &model
}

type DailyBootReasonView struct {
	Date     time.Time
	ViewData []item
}

func newDailyBootReasonView(date time.Time) *DailyBootReasonView {
	view := DailyBootReasonView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "内核panic重启开机", Value: 0},
		item{Key: "开门狗设置重启，超时重启开机，或者其它死机错误", Value: 0},
		item{Key: "正常使用时，启动失败发生系统切换", Value: 0})

	return &view
}
