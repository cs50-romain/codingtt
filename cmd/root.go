package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	_, _ = getTimers()
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

		// This commands can have 0 arguments
		if command == "help" {
			getHelp(fields)
		} else if command == "exit" {
			getExit(fields)
			break
		}

		if len(fields) <= 1 {
			continue
		}

		if command == "create" {
			getCreate(fields[1:])
		} else if command == "start" {
			getStart(fields[1:])
		} else if command == "stop" {
			getStop(fields[1:])
		} else if command == "pause" {
			getPause(fields[1:])
		} else if command == "restart" {
			getRestart(fields[1:])
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
