package mongomodel

import (
	"time"
)

func Item(key string, value int) item {
	return item{
		Key:   key,
		Value: value,
	}
}

type item struct {
	Key   string
	Value int
}

type BuildVersionMapping struct {
	Date              time.Time
	OnlineRobotCounts int
	BuildMapping      []int
}

func NewBuildVersionMapping(date time.Time) *BuildVersionMapping {
	return &BuildVersionMapping{
		Date:              date,
		OnlineRobotCounts: 0,
		BuildMapping:      make([]int, 0),
	}
}

type DailyVersionView struct {
	Date     time.Time
	ViewData []item
}

func NewDailyVersionView(date time.Time) *DailyVersionView {
	view := DailyVersionView{
		Date: date,
	}

	view.ViewData = append(view.ViewData,
		item{Key: "0000000000REL", Value: 0},
		item{Key: "2016071600REL", Value: 0},
		item{Key: "2016071501REL", Value: 0},
		item{Key: "2016072201REL", Value: 0},
		item{Key: "2016072500REL", Value: 0},
		item{Key: "2016080400REL", Value: 0},
		item{Key: "2016081200REL", Value: 0},
		item{Key: "2016081600REL", Value: 0},
		item{Key: "2016091200REL", Value: 0},
		item{Key: "2016091800REL", Value: 0},
		item{Key: "2016092600REL", Value: 0},
		item{Key: "2016092601REL", Value: 0},
		item{Key: "2016102100REL", Value: 0},
		item{Key: "2016102400REL", Value: 0},
		item{Key: "2016120200REL", Value: 0},
		item{Key: "2016120700REL", Value: 0},
		item{Key: "2017010301REL", Value: 0},
		item{Key: "2017011902REL", Value: 0},
		item{Key: "2017021002REL", Value: 0},
		item{Key: "2017021600REL", Value: 0},
		item{Key: "2017031600REL", Value: 0},
		item{Key: "2017041700REL", Value: 0},
		item{Key: "2017042701REL", Value: 0},
		item{Key: "2017042800REL", Value: 0},
		item{Key: "2017051501REL", Value: 0},
		item{Key: "2017052600REL", Value: 0},
		item{Key: "2017060800REL", Value: 0},
		item{Key: "2017062000REL", Value: 0},
		item{Key: "2017090700REL", Value: 0},
		item{Key: "2017110300REL", Value: 0},
		item{Key: "2017120602REL", Value: 0},
		item{Key: "2018013002REL", Value: 0})

	return &view
}

func NewRubysDailyVersionView(date time.Time) *DailyVersionView {
	view := DailyVersionView{
		Date: date,
	}

	view.ViewData = append(view.ViewData,
		item{Key: "0000000000REL", Value: 0},
		item{Key: "2017091802REL", Value: 0})

	return &view
}
