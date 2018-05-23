package mongomodel

import (
	"time"
)

type VoiceVersionMapping struct {
	Date              time.Time
	OnlineRobotCounts int
	VoiceMapping      []int
}

func NewVoiceVersionMapping(date time.Time) *VoiceVersionMapping {
	return &VoiceVersionMapping{
		Date:              date,
		OnlineRobotCounts: 0,
		VoiceMapping:      make([]int, 0),
	}
}

type RubysVoiceNode struct {
	SN    string
	Value int
}

type RubysVoiceVersionMapping struct {
	Date              time.Time
	OnlineRobotCounts int
	VoiceMapping      []RubysVoiceNode
}

func NewRubysVoiceVersionMapping(date time.Time) *RubysVoiceVersionMapping {
	return &RubysVoiceVersionMapping{
		Date:              date,
		OnlineRobotCounts: 0,
		VoiceMapping:      make([]RubysVoiceNode, 0),
	}
}

type VoiceVersionModel struct {
	View    *DailyVoiceVersionView
	Typemap map[int]int
}

func NewVoiceVersionModel(date time.Time) *VoiceVersionModel {
	model := VoiceVersionModel{
		View:    newDailyVoiceVersionView(date),
		Typemap: make(map[int]int),
	}

	model.Typemap[1] = 0
	model.Typemap[2] = 1
	model.Typemap[3] = 2
	model.Typemap[1005] = 3
	model.Typemap[1004] = 4
	model.Typemap[1003] = 5
	model.Typemap[1002] = 6
	model.Typemap[1001] = 7
	model.Typemap[1000] = 8

	return &model
}

type DailyVoiceVersionView struct {
	Date     time.Time
	ViewData []item
}

func newDailyVoiceVersionView(date time.Time) *DailyVoiceVersionView {
	view := DailyVoiceVersionView{
		Date: date,
	}
	view.ViewData = append(view.ViewData,
		item{Key: "标准普通话版", Value: 0},
		item{Key: "台湾普通话版", Value: 0},
		item{Key: "英文版", Value: 0},
		item{Key: "机器人版", Value: 0},
		item{Key: "动漫儿童版", Value: 0},
		item{Key: "后宫嫔妃版", Value: 0},
		item{Key: "萌妹子版", Value: 0},
		item{Key: "播音员版", Value: 0},
		item{Key: "粤语版", Value: 0})

	return &view
}
