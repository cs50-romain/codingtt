package cmd

import (
	"fmt"
	"log"
	"time"
)

type Command struct {
	Name		string
	Desc		string
	//Arguments	[]string -> Can be used to hold any sort of arguments for a command
	CallBack	func([]string)
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
			Desc: "Starts the timer. start <timer name>",
			CallBack: getStart,
		},
		"stop": {
			Name: "stop",
			Desc: "Stops the timer. Does not exit the program.",
			CallBack: getStop,
		},
		"pause": {
			Name: "pause",
			Desc: "Pauses the timer. pause <timer name>.",
			CallBack: getPause,
		},
		"restart": {
			Name: "restart",
			Desc: "Restarts the timer. restart <timer name>",
			CallBack: getRestart,
		},
	}
}

func getHelp(args []string) {
	availableCommands := getCommands()

	if len(args) <= 1 {
		for _, cmd := range availableCommands {
			fmt.Printf(" -%s: %s\n", cmd.Name, cmd.Desc)
		}
	} else {
		cmdName := args[1]
		cmd := availableCommands[cmdName]
		fmt.Printf(" -%s: %s\n", cmd.Name, cmd.Desc)
	}
}

func getExit(args []string) {
	timerss := getTimers()
	for _, timer := range timerss {
		if timer.Export != true || timer.Name == "unknown" {
			continue
		}
		err := timer.ExportToCsv(CSV_FILE)
		if err != nil {
			log.Printf("error export data for timer: %s\n", timer.Name)
		}
	}
}

func getStart(args []string) {
	timerName := parseName(args)
	
	var exportOpt bool
	if len(args) < 2 {
		exportOpt = true
	} else if args[1] == "false" {
		exportOpt = false
	} else {
		exportOpt = true
	}

	timer := CreateTimer(timerName, exportOpt)
	timers[timerName] = timer

	Stack.Push(timer)

	timer.Start = time.Now()
	fmt.Printf("Started %s!\n", timer.Name)
}

func getStop(args []string) {
	var timer *Timer

	timerName := parseName(args)
	/*
	if timerExists(timerName) == false {
		fmt.Println("Invalid timer name")
		return
	}
	*/
	if len(args) == 0 {
		timer = Stack.Pop()
	} else {
		timer = timers[timerName]
	}

	if timer.Start.IsZero() {
		fmt.Println("Timer has not been started.")
		return
	}

	timer.Stop = time.Now()
	fmt.Println("Stopping ", timer.Name)

	timer.Total = timer.CalcTotalSeconds(timer.Start, timer.Stop)
	fmt.Printf("Time spent coding: %s\n", timer.formatTotalTime(timer.Total))
}

func getPause(args []string) {
	timerName := parseName(args)

	var timer *Timer
	if timerName == "unknown" {
		timer = Stack.Peek()
	} else {
		timer = timers[timerName]
	}

	if timer.Start.IsZero() {
		fmt.Println("Timer has not been started.")
		return
	}

	timer.Pause[0] = time.Now()
	fmt.Println("pausing", timer.Name)
}

func getRestart(args []string) {
	timerName := parseName(args)

	var timer *Timer
	if timerName == "unknown" {
		timer = Stack.Peek()
		timer.Pause[1] = time.Now()
	} else {
		timer = timers[timerName]
	}

	if timer.Pause[0].IsZero() {
		fmt.Println("Timer has not been paused.")
		return
	}

	timer.Pause[1] = time.Now()
	totalPauseTime := timer.CalcTotalSeconds(timer.Pause[0], timer.Pause[1])
	timer.Total = timer.Total - totalPauseTime
	fmt.Println("Restarting")
}

func parseName(args []string) string {
	if len(args) == 0 {
		return "unknown"
	} else {
		return args[0]
	}
}
