package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Belstowe/distrib-pc/lab2/dsp"
	"github.com/Belstowe/distrib-pc/lab2/tasks"
)

var (
	records []tasks.Task
)

func init() {
	readRecords()
	sortRecords()
}

func main() {
	fmt.Println(records)
	fmt.Printf("NFDH: %v\n", assert(dsp.NFDH(records, 10)))
	fmt.Printf("FFDH: %v\n", assert(dsp.FFDH(records, 10)))
}

func assert[T any](res T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func readRecords() {
	f := assert(os.Open("test_tasks.csv"))

	r := csv.NewReader(f)
	r.Comma = ' '
	r.Comment = '#'

	rawRecords := assert(r.ReadAll())
	for _, rawRecord := range rawRecords {
		records = append(records, tasks.Task{
			R: assert(strconv.Atoi(rawRecord[0])),
			T: assert(strconv.Atoi(rawRecord[1])),
		})
	}
}

func sortRecords() {
	min, max := records[0].T, records[1].T
	for _, record := range records {
		if record.T < min {
			min = record.T
		} else if record.T > max {
			max = record.T
		}
	}

	recordBuckets := make([][]tasks.Task, max-min+1)
	for i := 0; i < max-min+1; i++ {
		recordBuckets[i] = make([]tasks.Task, 0)
	}

	for _, record := range records {
		recordBuckets[record.T-min] = append(recordBuckets[record.T-min], record)
	}

	records = records[:0]
	for i := len(recordBuckets) - 1; i >= 0; i-- {
		records = append(records, recordBuckets[i]...)
	}
}
