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
		"create": {
			Name: "create",
			Desc: "Create a new timer. create <name> <export true/false>",
			CallBack: getCreate,
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
	// Export data to csv
	// err := util.Export(timer)
	timerss, err := getTimers()
	if err != nil {
		fmt.Println(err)
	}

	for _, timer := range timerss {
		fmt.Println(timer)
		if timer.Export != true {
			continue
		}
		err := timer.exportToCsv()
		if err != nil {
			log.Printf("error export data for timer: %s\n", timer.Name)
		}
	}
}

func getCreate(args []string) {
	var exportOption bool
	timerName := args[0]
	
	// Check if timer exists in timers
	if _, ok := timers[timerName]; ok {
		fmt.Println("timer already exists!")
		return
	}

	if len(args) <= 1 {
		exportOption = true
		timer := CreateTimer(timerName, exportOption)
		timers[timerName] = timer
		return 
	}

	if args[1] == "false" {
		exportOption = false
	} else {
		exportOption = true
	}
	timer := CreateTimer(timerName, exportOption)
	timers[timerName] = timer
}

func getStart(args []string) {
	timerName := args[0]
	if timerExists(timerName) == false {
		fmt.Println("Invalid timer name")
		return
	}


	timer := timers[timerName]
	timer.Start = time.Now()
	fmt.Println("Started timer, program mf!")
}

func getStop(args []string) {
	timerName := args[0]
	if timerExists(timerName) == false {
		fmt.Println("Invalid timer name")
		return
	}

	timer := timers[timerName]

	timer.Stop = time.Now()
	fmt.Println("Stopping")

	timer.Total = timer.calcTotalSeconds(timer.Start, timer.Stop)
	fmt.Printf("Time spent coding: %s\n", timer.formatTotalTime(timer.Total))
}

func getPause(args []string) {
	timerName := args[0]
	if timerExists(timerName) == false {
		fmt.Println("Invalid timer name")
		return
	}

	timer := timers[timerName]

	timer.Pause[0] = time.Now()
	fmt.Println("pausing")
}

func getRestart(args []string) {
	timerName := args[0]
	if timerExists(timerName) == false {
		fmt.Println("Invalid timer name")
		return
	}

	timer := timers[timerName]

	timer.Pause[1] = time.Now()
	totalPauseTime := timer.calcTotalSeconds(timer.Pause[0], timer.Pause[1])
	timer.Total = timer.Total - totalPauseTime
	fmt.Println("Restarting")
}
