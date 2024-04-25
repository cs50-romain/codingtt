package main

import (
	"cs50-romain/codingtt/cmd"
	"log"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
