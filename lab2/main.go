package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Belstowe/distrib-pc/lab2/dsp"
	"github.com/Belstowe/distrib-pc/lab2/tasks"
)

const (
	None int = iota
	NFDH
	FFDH
)

func main() {
	filepath, algo, rn := flagParse()
	records := readRecords(filepath)
	records = sortRecords(records)

	switch algo {
	case None:
		fmt.Println(records)
	case NFDH:
		fmt.Println(assert(dsp.NFDH(records, rn)))
	case FFDH:
		fmt.Println(assert(dsp.FFDH(records, rn)))
	}
}

func assert[T any](res T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func invalidUsage(format string, v ...any) {
	log.Printf(format, v...)
	flag.Usage()
	os.Exit(1)
}

func flagParse() (string, int, int) {
	filepath := flag.String("f", "", "path to file with tasks (.csv format, split with spaces, machines & time columns)")
	nfdhFlag := flag.Bool("nfdh", false, "use NFDH (Next Fit Decreasing Height) algorithm")
	ffdhFlag := flag.Bool("ffdh", false, "use FFDH (First First Decreasing Height) algorithm")
	rn := flag.Int("n", 0, "num of elementary machines")

	flag.Parse()

	if *filepath == "" {
		invalidUsage("path should be set")
	}

	algo := None
	if *nfdhFlag {
		algo = NFDH
	} else if *ffdhFlag {
		algo = FFDH
	} else {
		invalidUsage("one of algorithm flag set required")
	}

	if *rn <= 0 {
		invalidUsage("num of elementary machines either not set or is not positive")
	}

	return *filepath, algo, *rn
}

func readRecords(filepath string) []tasks.Task {
	f := assert(os.Open("test_tasks.csv"))

	r := csv.NewReader(f)
	r.Comma = ' '
	r.Comment = '#'

	records := make([]tasks.Task, 0)
	rawRecords := assert(r.ReadAll())
	for _, rawRecord := range rawRecords {
		records = append(records, tasks.Task{
			R: assert(strconv.Atoi(rawRecord[0])),
			T: assert(strconv.Atoi(rawRecord[1])),
		})
	}
	return records
}

func sortRecords(records []tasks.Task) []tasks.Task {
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

	sortedRecords := make([]tasks.Task, 0, len(records))
	for i := len(recordBuckets) - 1; i >= 0; i-- {
		sortedRecords = append(sortedRecords, recordBuckets[i]...)
	}
	return sortedRecords
}
