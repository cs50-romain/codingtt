package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestCalcTotalSeconds(t *testing.T) {
	var tests = []struct {
		startTime, stopTime time.Time
		want int
	}{
		{time.Date(2009, time.November, 10, 23, 1, 34, 0, time.UTC), time.Date(2009, time.November, 10, 23, 2, 23, 0, time.UTC), 49},
		{time.Date(2009, time.November, 10, 10, 5, 0, 0, time.UTC), time.Date(2009, time.November, 10, 15, 11, 25, 0, time.UTC), 18385},
	}

	var timer Timer
	for _, test := range tests {
		testname := fmt.Sprintf("%v,%v", test.startTime, test.stopTime)
		t.Run(testname, func(t *testing.T) {
			result := timer.CalcTotalSeconds(test.startTime, test.stopTime)
			if result != test.want {
				t.Errorf("got %d, want %d", result, test.want)
			}
		})
	}
}

func TestExportToCsv(t *testing.T) {
	timer := &Timer{
		Name: "example",
		Export: true,
		Start: time.Date(2009, time.November, 10, 23, 1, 34, 0, time.UTC),
		Stop:time.Date(2009, time.November, 10, 23, 38, 29, 0, time.UTC),
		Total: 8938,
	}

	if err := timer.ExportToCsv("example.csv"); err != nil {
		t.Errorf("error: unable to export")
	}
}
