package cmd

import (
	"fmt"
	"log"
	"time"
)

type Command struct {
	Name	string
	Desc	string
	CallBack func()
}

func getCommands() map[string]Command {
	return map[string]Command {
		"help": {
			Name: "help",
			Desc: "Shows commands and their description",
			CallBack: getHelp,
		},
		"exit": {
			Name: "exit",
			Desc: "Exits the program",
			CallBack: getExit,
		},
		"start": {
			Name: "start",
			Desc: "Starts the timer",
			CallBack: getStart,
		},
		"stop": {
			Name: "stop",
			Desc: "Stops the timer. Does not exit the program.",
			CallBack: getStop,
		},
		"pause": {
			Name: "pause",
			Desc: "Pauses the timer.",
			CallBack: getPause,
		},
		"restart": {
			Name: "restart",
			Desc: "Restarts the timer.",
			CallBack: getRestart,
		},
	}
}

func getHelp() {
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" -%s: %s\n", cmd.Name, cmd.Desc)
	}
}

func getExit() {
	// Export data to csv
	// err := util.Export(timer)
	err := timer.ExportToCsv()
	if err != nil {
		log.Panic(err)
	}
}

func getStart() {
	timer.Start = time.Now()
	fmt.Println("Started timer, program mf!")
}

func getStop() {
	timer.Stop = time.Now()
	fmt.Println("Stopping")

	timer.Total = timer.calcTotalSeconds(timer.Start, timer.Stop)
	fmt.Printf("Time spent coding: %s\n", timer.formatTotalTime(timer.Total))
}

func getPause() {
	timer.Pause[0] = time.Now()
	fmt.Println("pausing")
}

func getRestart() {
	timer.Pause[1] = time.Now()
	totalPauseTime := timer.calcTotalSeconds(timer.Pause[0], timer.Pause[1])
	timer.Total = timer.Total - totalPauseTime
	fmt.Println("Restarting")
}
