package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)	

func Start() error {
	// Start repl
	if err := startRepl(); err != nil {
		return err
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

		fields := parse(input)
		if len(fields) == 0 {
			continue
		}
		command := fields[0]
		var args []string

		if len(fields) > 1 {
			args = fields[1:]
		}

		if command == "help" {
			getHelp(args)
		} else if command == "exit" {
			getExit(args)
			break
		} else if command == "start" {
			getStart(args)
		} else if command == "stop" {
			getStop(args)
		} else if command == "pause" {
			getPause(args)
		} else if command == "restart" {
			getRestart(args)
		} else {
			fmt.Printf("Invalid command. Type help to show available commands\n")
			continue
		}
	}
	return nil
}

// Testable and should be tested
func parse(in string) ([]string) {
	in = strings.ToLower(in)
	fields := strings.Fields(in)
	return fields
}

func fileIsEmpty(filename string) bool {
	info,_ := os.Stat(filename)
	if info.Size() == 0 {
		return true
	}
	return false
}
