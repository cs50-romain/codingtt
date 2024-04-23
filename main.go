package main

import (
	"cs50-romain/codingtt/cmd"
	"log"
	"os"
)

func main() {
	options := os.Args

	if err := cmd.Start(options); err != nil {
		log.Fatal(err)
	}
}
