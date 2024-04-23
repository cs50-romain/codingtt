package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Timer struct {
	Start time.Time
	Stop time.Time
	Pause [2]time.Time
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
		go showTimer()
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
			hour, mins, sec := timer.calculateTime()
			fmt.Printf("Time spent coding: %d:%d:%d hours\n", hour, mins, sec)
			// Exit
			break
		} else if command == "pause" {
			timer.Pause[0] = time.Now()
			fmt.Println("pausing")
		} else if command == "start" {
			timer.Pause[1] = time.Now()
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

// Figure this out later. Show a timer.
func showTimer() {
	//for {
		//fmt.Print(time.Now())
	//}
}

func (t *Timer) calculateTime() (int, int, int) {
	fmt.Println(timer.Stop.Clock())
	fmt.Println(timer.Start.Clock())
	hours := math.Abs(float64(timer.Stop.Hour() - timer.Start.Hour()))
	mins := math.Abs(float64(timer.Stop.Minute() - timer.Start.Minute()))
	sec := math.Abs(float64(timer.Stop.Second() - timer.Start.Second()))
	return int(hours), int(mins), int(sec)
}
