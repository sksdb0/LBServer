package mongomodel

import (
	"time"
)

type DailyNewRobotView struct {
	Date             time.Time
	RobotCounts      int
	Robot13871Counts int
	Robot15339Counts int
}

func NewDailyNewRobotView(date time.Time) *DailyNewRobotView {
	view := DailyNewRobotView{
		Date: date,
	}

	return &view
}
