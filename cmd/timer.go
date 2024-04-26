package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const HOURS_TO_SECONDS = 3600
const MINUTES_TO_SECONDS = 60
const CSV_FILE = "data.csv"

type Timer struct {
	Start time.Time
	Stop time.Time
	Pause [2]time.Time
	Total int
}

func (t *Timer) calcTotalSeconds(startTime, stopTime time.Time) (int) {
	var startTotalSeconds, stopTotalSeconds, totalSeconds int

	startTotalSeconds = (startTime.Hour() * HOURS_TO_SECONDS) + (startTime.Minute() * MINUTES_TO_SECONDS) + startTime.Second()

	stopTotalSeconds = (stopTime.Hour() * HOURS_TO_SECONDS) + (stopTime.Minute() * MINUTES_TO_SECONDS) + stopTime.Second()

	totalSeconds = stopTotalSeconds - startTotalSeconds

	//fmt.Printf("Time start: %d:%d:%d; Time end: %d:%d:%d, totalSeconds = %d\n", t.Start.Hour(), t.Start.Minute(), t.Start.Second(), t.Stop.Hour(), t.Stop.Minute(), t.Stop.Second(), totalSeconds)

	return totalSeconds
}

func (t *Timer) formatTotalTime(totalSeconds int) string {
	var hours, mins, secs int
	for totalSeconds >= 3600 {
		totalSeconds -= 3600
		hours++
	}

	for totalSeconds >= 60 {
		totalSeconds -= 60
		mins++
	}

	secs = totalSeconds

	return fmt.Sprintf("%02d:%02d:%02d", int(hours), int(mins), int(secs))
}

func (t *Timer) ExportToCsv() error {
	var data [][]string

	file, err := os.OpenFile(CSV_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%v/%d/%d", month, day, year)
	header := []string{"Date", "Coding Time", "Notes"}
	totalStr := fmt.Sprintf("%s", t.formatTotalTime(t.Total))
	line := []string{date, totalStr, "notes"}

	if fileIsEmpty(CSV_FILE) {
		data = [][]string{
			header,
			line,
		}
	} else {
		data = [][]string{
			line,
		}
	}

	w := csv.NewWriter(file)
	if err = w.WriteAll(data); err != nil {
		return err
	}

	w.Flush()
	return nil	
}
