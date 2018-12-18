package main

import (
	"lemon-robot-cli/events"
	"lemon-robot-cli/logger"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		logger.Error("Please tell me what you want to do, for example:")
		logger.Error("lr build")
		logger.Error("lr run")
		os.Exit(-1)
	}
	events.Build()
}
