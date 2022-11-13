package dsp

import "github.com/Belstowe/distrib-pc/lab2/tasks"

func TotalTime(taskLevels [][]tasks.Task) int {
	tt := 0
	for _, taskLevel := range taskLevels {
		tt += taskLevel[0].T
	}
	return tt
}

func TimeBound(taskLevels [][]tasks.Task) float64 {
	tavg := 0
	taskNum := 0
	for _, taskLevel := range taskLevels {
		for _, task := range taskLevel {
			tavg += task.T * task.R
			taskNum++
		}
	}
	return float64(tavg) / float64(taskNum)
}

func FnDeviation(taskLevels [][]tasks.Task) float64 {
	return (float64(TotalTime(taskLevels)) - TimeBound(taskLevels)) / TimeBound(taskLevels)
}
