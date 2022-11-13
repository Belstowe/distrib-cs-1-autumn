package dsp

import (
	"fmt"

	"github.com/Belstowe/distrib-pc/lab2/tasks"
)

func FFDH(records []tasks.Task, n int) ([][]tasks.Task, error) {
	if records[0].R > n {
		return nil, fmt.Errorf("only having %d processes, while the biggest task requires %d processes", n, records[0].R)
	}

	taskLevels := make([][]tasks.Task, 0)
	levelTakenProc := make([]int, 0)
outerLoop:
	for _, task := range records {
		for i := range taskLevels {
			if levelTakenProc[i]+task.R <= n {
				levelTakenProc[i] += task.R
				taskLevels[i] = append(taskLevels[i], task)
				continue outerLoop
			}
		}
		taskLevels = append(taskLevels, []tasks.Task{task})
		levelTakenProc = append(levelTakenProc, task.R)
	}

	return taskLevels, nil
}
