package cmd

import (
	"bufio"
	"errors"
	"fmt"
	//"math"
	"os"
	"strings"
	"time"
)

const HOURS_TO_SECONDS = 3600
const MINUTES_TO_SECONDS = 60

type Timer struct {
	Start time.Time
	Stop time.Time
	Pause [2]time.Time
}


func (t *Timer) calcTotalSeconds() (int) {
	var startTotalSeconds, stopTotalSeconds, totalSeconds int

	startTotalSeconds = (t.Start.Hour() * HOURS_TO_SECONDS) + (t.Start.Minute() * MINUTES_TO_SECONDS) + t.Start.Second()

	stopTotalSeconds = (t.Stop.Hour() * HOURS_TO_SECONDS) + (t.Stop.Minute() * MINUTES_TO_SECONDS) + t.Stop.Second()

	totalSeconds = stopTotalSeconds - startTotalSeconds

	fmt.Printf("Time start: %d:%d:%d; Time end: %d:%d:%d, totalSeconds = %d\n", t.Start.Hour(), t.Start.Minute(), t.Start.Second(), t.Stop.Hour(), t.Stop.Minute(), t.Stop.Second(), totalSeconds)

	return totalSeconds

}

func (t *Timer) formatTotalTime(totalSeconds int) (int, int, int) {
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

	return int(hours), int(mins), int(secs)
}

var timer Timer

func Start(options []string) error {
	// Start repl
	//reader := bufio.NewReader(os.Stdin)
	
	if len(options) > 2 {
		return errors.New("Error: Too many arguments")
	}

	option := options[1]

	// Parsing args
	if option == "-start" {
		timer.Start = time.Now()
		fmt.Println("Timer started!")
		if err := startRepl(); err != nil {
			return err
		}
	} else if option == "-total" {
		// Show total amount of coding since first run of this timer
		// Import data from .csv
		// total, err := util.GetTotal()
	} else {
		return errors.New("Error: Invalid argument. Run help if needed.")
	}
	
	return nil
}

func startRepl() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		command, err := parse(input)
		if err != nil {
			return err
		}

		if command == "stop" {
			timer.Stop = time.Now()
			fmt.Println("Stopping")
			// Calculate amount of time between start and stop and substart pause time
			// Export to .csv file
			// err := util.Export(timer)
			totalSeconds := timer.calcTotalSeconds()
			hour, mins, sec := timer.formatTotalTime(totalSeconds)
			fmt.Printf("Time spent coding: %d:%d:%d\n", hour, mins, sec)
			// Exit
			break
		} else if command == "pause" {
			timer.Pause[0] = time.Now()
			fmt.Println("pausing")
		} else if command == "start" {
			timer.Pause[1] = time.Now()
			
			// Calculate the total seconds since the first pause.
			// Add that total somewhere so it is remembered.

			fmt.Println("Restarting")
		} else if command == "exit" {
			fmt.Println("Exiting")
			os.Exit(0)
		} else { // command == total
			// Calculate total amount of coding time since timer.Start
		}
	}
	return nil
}

func parse(in string) (string, error) {
	commands := strings.Split(in, " ")
	return strings.TrimSuffix(commands[0], "\n"), nil
}
