package mongomodel

import (
	"time"
)

type CleanMapping struct {
	Date             time.Time
	CleanCounts      int
	CleanRobotCounts int
	Mapping          []int
}

func NewCleanMapping(date time.Time, length int) *CleanMapping {
	return &CleanMapping{
		Date:             date,
		CleanCounts:      0,
		CleanRobotCounts: 0,
		Mapping:          make([]int, length),
	}
}

type RubysCleanNode struct {
	SN        string
	CleanInfo int
}

type RubysCleanMapping struct {
	Date             time.Time
	CleanCounts      int
	CleanRobotCounts int
	Mapping          []RubysCleanNode
}

func NewRubysCleanMapping(date time.Time) *RubysCleanMapping {
	return &RubysCleanMapping{
		Date:             date,
		CleanCounts:      0,
		CleanRobotCounts: 0,
	}
}

type DailyCleanView struct {
	Date              time.Time
	CleanCounts       int
	CleanRobotCounts  int
	OnlineRobotCounts int
}

func NewDailyCleanView(date time.Time) *DailyCleanView {
	return &DailyCleanView{
		Date: date,
	}
}

type WeeklyCleanView struct {
	Date              time.Time
	CleanCounts       int
	CleanRobotCounts  int
	OnlineRobotCounts int
	ViewDataDays      []item
	ViewDataSchedule  []item
}

func NewWeeklyView(date time.Time) *WeeklyCleanView {
	view := WeeklyCleanView{
		Date: date,
	}

	view.ViewDataDays = append(view.ViewDataDays,
		item{Key: "1天", Value: 0},
		item{Key: "2天", Value: 0},
		item{Key: "3天", Value: 0},
		item{Key: "4天", Value: 0},
		item{Key: "5天", Value: 0},
		item{Key: "6天", Value: 0},
		item{Key: "7天", Value: 0})

	return &view
}

type monthlyView struct {
	Date              time.Time
	CleanCounts       int
	CleanRobotCounts  int
	OnlineRobotCounts int
}

func NewMonthlyView(date time.Time) *monthlyView {
	return &monthlyView{
		Date: date,
	}
}
