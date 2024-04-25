package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var timer Timer

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

		if command == "help" {
			getHelp()
		} else if command == "start" {
			getStart()
		} else if command == "stop" {
			getStop()
		} else if command == "pause" {
			getPause()
		} else if command == "restart" {
			getRestart()
		} else if command == "exit" {
			getExit()
			break
		} else {
			fmt.Printf("Invalid command. Type help to show available commands\n")
			continue
		}
	}
	return nil
}

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
