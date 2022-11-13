package dsp

import (
	"fmt"

	"github.com/Belstowe/distrib-pc/lab2/tasks"
)

func NFDH(records []tasks.Task, n int) ([][]tasks.Task, error) {
	if records[0].R > n {
		return nil, fmt.Errorf("only having %d processes, while the biggest task requires %d processes", n, records[0].R)
	}

	tasksLength := len(records)
	taskLevels := make([][]tasks.Task, 0)
	if tasksLength == 0 {
		return taskLevels, nil
	}

	j := 0
outerLoop:
	for i := 0; ; i++ {
		takenProc := 0
		taskLevels = append(taskLevels, make([]tasks.Task, 0))
		for {
			takenProc += records[j].R
			if takenProc > n {
				continue outerLoop
			}
			taskLevels[i] = append(taskLevels[i], records[j])
			j++
			if j >= tasksLength {
				break outerLoop
			}
		}
	}

	return taskLevels, nil
}
