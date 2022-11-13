package output

import "github.com/Belstowe/distrib-pc/lab2/tasks"

type AlgoEfficiency struct {
	PerformanceTime   float64
	ScheduleTotalTime int
	ScheduleDeviation float64
	Schedule          [][]tasks.Task
}
